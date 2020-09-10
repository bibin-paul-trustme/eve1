// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// dnsmasq configlets for overlay and underlay interfaces towards domU

package zedrouter

import (
	"bufio"
	"fmt"
	"github.com/lf-edge/eve/pkg/pillar/base"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lf-edge/eve/pkg/pillar/types"
)

const dnsmasqStatic = `
# Automatically generated by zedrouter
except-interface=lo
bind-interfaces
# log-queries XXX way to noisy
# log-dhcp
quiet-dhcp
quiet-dhcp6
no-hosts
no-ping
bogus-priv
stop-dns-rebind
rebind-localhost-ok
neg-ttl=10
dhcp-ttl=600
`

// dnsmasqLeaseDir is used to for the leases
// of bridgeNames
var dnsmasqLeaseDir = runDirname + "/dnsmasq.leases/"

// dnsmasqLeasePath provides a unique file
// We traverse the dnsmasqLeaseDir directory to get the list of bridgeNames
func dnsmasqLeasePath(bridgeName string) string {
	leasePathname := dnsmasqLeaseDir + "/" + bridgeName
	return leasePathname
}

func dnsmasqBridgeNames() []string {
	var bridgeNames []string

	locations, err := ioutil.ReadDir(dnsmasqLeaseDir)
	if err != nil {
		log.Error(err)
		return bridgeNames
	}

	for _, location := range locations {
		bridgeNames = append(bridgeNames, location.Name())
	}
	return bridgeNames
}

func dnsmasqInitDirs() {
	if _, err := os.Stat(dnsmasqLeaseDir); err != nil {
		log.Infof("Create %s\n", dnsmasqLeaseDir)
		if err := os.Mkdir(dnsmasqLeaseDir, 0755); err != nil {
			log.Fatal(err)
		}
	} else {
		// dnsmasq needs to read as nobody
		if err := os.Chmod(dnsmasqLeaseDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func dnsmasqConfigFile(bridgeName string) string {
	cfgFilename := "dnsmasq." + bridgeName + ".conf"
	return cfgFilename
}

func dnsmasqConfigPath(bridgeName string) string {
	cfgFilename := dnsmasqConfigFile(bridgeName)
	cfgPathname := runDirname + "/" + cfgFilename
	return cfgPathname
}

func dnsmasqDhcpHostDir(bridgeName string) string {
	dhcphostsDir := runDirname + "/dhcp-hosts." + bridgeName
	return dhcphostsDir
}

// createDnsmasqConfiglet
// When we create a linux bridge we set this up
// Also called when we need to update the ipsets
func createDnsmasqConfiglet(
	bridgeName string, bridgeIPAddr string,
	netconf *types.NetworkInstanceConfig, hostsDir string,
	ipsets []string, Ipv4Eid bool, uplink string, dnsServers []net.IP) {

	log.Infof("createDnsmasqConfiglet(%s, %s) netconf %v, ipsets %v uplink %s dnsServers %v",
		bridgeName, bridgeIPAddr, netconf, ipsets, uplink, dnsServers)

	cfgPathname := dnsmasqConfigPath(bridgeName)
	// Delete if it exists
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("createDnsmasqConfiglet %v",
				err)
			log.Errorln(errStr)
		}
	}
	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("createDnsmasqConfiglet failed ", err)
	}
	defer file.Close()

	// Create a dhcp-hosts directory to be used when hosts are added
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)

	file.WriteString(dnsmasqStatic)
	file.WriteString(fmt.Sprintf("dhcp-leasefile=%s\n",
		dnsmasqLeasePath(bridgeName)))

	// Pick file where dnsmasq should send DNS read upstream
	// If we have no uplink for this network instance that is nowhere
	// If we have an uplink but no dnsServers for it, then we let
	// dnsmasq use the host's /etc/resolv.conf
	if uplink == "" {
		file.WriteString("no-resolv\n")
	} else if len(dnsServers) != 0 {
		for _, s := range dnsServers {
			file.WriteString(fmt.Sprintf("server=%s@%s\n", s, uplink))
		}
		file.WriteString("no-resolv\n")
	}

	for _, ipset := range ipsets {
		file.WriteString(fmt.Sprintf("ipset=/%s/ipv4.%s,ipv6.%s\n",
			ipset, ipset, ipset))
	}
	file.WriteString(fmt.Sprintf("pid-file=/var/run/dnsmasq.%s.pid\n",
		bridgeName))
	file.WriteString(fmt.Sprintf("interface=%s\n", bridgeName))
	isIPv6 := false
	if bridgeIPAddr != "" {
		ip := net.ParseIP(bridgeIPAddr)
		if ip == nil {
			log.Fatalf("createDnsmasqConfiglet failed to parse IP %s",
				bridgeIPAddr)
		}
		isIPv6 = (ip.To4() == nil)
		file.WriteString(fmt.Sprintf("listen-address=%s\n",
			bridgeIPAddr))
	} else {
		// XXX error if there is no bridgeIPAddr?
	}
	file.WriteString(fmt.Sprintf("hostsdir=%s\n", hostsDir))
	file.WriteString(fmt.Sprintf("dhcp-hostsdir=%s\n", dhcphostsDir))

	ipv4Netmask := "255.255.255.0" // Default unless there is a Subnet
	dhcpRange := bridgeIPAddr      // Default unless there is a DhcpRange

	// By default dnsmasq advertizes a router (and we can have a
	// static router defined in the NetworkInstanceConfig).
	// To support airgap networks we interpret gateway=0.0.0.0
	// to not advertize ourselves as a router. Also,
	// if there is not an explicit dns server we skip
	// advertising that as well.
	advertizeRouter := true
	var router string

	if netconf.Logicallabel == "" {
		log.Infof("Internal switch without external port case, dnsmasq suppress router advertize\n")
		advertizeRouter = false
	} else if Ipv4Eid {
		advertizeRouter = false
	} else if netconf.Gateway != nil {
		if netconf.Gateway.IsUnspecified() {
			advertizeRouter = false
		} else {
			router = netconf.Gateway.String()
		}
	} else if bridgeIPAddr != "" {
		router = bridgeIPAddr
	} else {
		advertizeRouter = false
	}
	if netconf.DomainName != "" {
		if isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:domain-search,%s\n",
				netconf.DomainName))
		} else {
			file.WriteString(fmt.Sprintf("dhcp-option=option:domain-name,%s\n",
				netconf.DomainName))
		}
	}
	advertizeDns := false
	if Ipv4Eid {
		advertizeDns = true
	}
	for _, ns := range netconf.DnsServers {
		advertizeDns = true
		file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server,%s\n",
			ns.String()))
	}
	if netconf.NtpServer != nil {
		file.WriteString(fmt.Sprintf("dhcp-option=option:ntp-server,%s\n",
			netconf.NtpServer.String()))
	}
	if netconf.Subnet.IP != nil {
		ipv4Netmask = net.IP(netconf.Subnet.Mask).String()
	}
	if netconf.Subnet.IP != nil {
		if advertizeRouter {
			// Network prefix "255.255.255.255" will force packets to go through
			// dom0 virtual router that makes the packets pass through ACLs and flow log.
			file.WriteString(fmt.Sprintf("dhcp-option=option:netmask,%s\n",
				"255.255.255.255"))
		} else {
			file.WriteString(fmt.Sprintf("dhcp-option=option:netmask,%s\n",
				ipv4Netmask))
		}
	}
	if advertizeRouter {
		// IPv6 XXX needs to be handled in radvd
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router,%s\n",
				router))
			file.WriteString(fmt.Sprintf("dhcp-option=option:classless-static-route,%s/32,%s,%s,%s,%s,%s\n",
				router, "0.0.0.0",
				"0.0.0.0/0", router,
				netconf.Subnet.String(), router))
		}
	} else {
		log.Infof("createDnsmasqConfiglet: no router\n")
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router\n"))
		}
		if !advertizeDns {
			// Handle isolated network by making sure
			// we are not a DNS server. Can be overridden
			// with the DnsServers above
			log.Infof("createDnsmasqConfiglet: no DNS server\n")
			file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server\n"))
		}
	}
	if netconf.DhcpRange.Start != nil {
		dhcpRange = netconf.DhcpRange.Start.String()
	}
	if isIPv6 {
		file.WriteString(fmt.Sprintf("dhcp-range=::,static,0,60m\n"))
	} else {
		file.WriteString(fmt.Sprintf("dhcp-range=%s,static,%s,60m\n",
			dhcpRange, ipv4Netmask))
	}
}

func addhostDnsmasq(bridgeName string, appMac string, appIPAddr string,
	hostname string) {

	log.Infof("addhostDnsmasq(%s, %s, %s, %s)\n", bridgeName, appMac,
		appIPAddr, hostname)
	ip := net.ParseIP(appIPAddr)
	if ip == nil {
		log.Fatalf("addhostDnsmasq failed to parse IP %s", appIPAddr)
	}
	isIPv6 := (ip.To4() == nil)
	suffix := ".inet"
	if isIPv6 {
		suffix += "6"
	}

	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)
	cfgPathname := dhcphostsDir + "/" + appMac + suffix

	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("addhostDnsmasq failed ", err)
	}
	defer file.Close()
	if isIPv6 {
		file.WriteString(fmt.Sprintf("%s,[%s],%s\n",
			appMac, appIPAddr, hostname))
	} else {
		file.WriteString(fmt.Sprintf("%s,id:*,%s,%s\n",
			appMac, appIPAddr, hostname))
	}
	file.Close()
}

func removehostDnsmasq(bridgeName string, appMac string, appIPAddr string) {

	log.Infof("removehostDnsmasq(%s, %s, %s)\n",
		bridgeName, appMac, appIPAddr)
	ip := net.ParseIP(appIPAddr)
	if ip == nil {
		log.Fatalf("removehostDnsmasq failed to parse IP %s", appIPAddr)
	}
	isIPv6 := (ip.To4() == nil)
	suffix := ".inet"
	if isIPv6 {
		suffix += "6"
	}

	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)

	cfgPathname := dhcphostsDir + "/" + appMac + suffix
	if _, err := os.Stat(cfgPathname); err != nil {
		log.Infof("removehostDnsmasq(%s, %s) failed: %s\n",
			bridgeName, appMac, err)
	} else {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("removehostDnsmasq %v", err)
			log.Errorln(errStr)
		}
	}
}

func deleteDnsmasqConfiglet(bridgeName string) {

	log.Infof("deleteDnsmasqConfiglet(%s)\n", bridgeName)
	cfgPathname := dnsmasqConfigPath(bridgeName)
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("deleteDnsmasqConfiglet %v",
				err)
			log.Errorln(errStr)
		}
	}
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)
	if err := RemoveDirContent(dhcphostsDir); err != nil {
		errStr := fmt.Sprintf("deleteDnsmasqConfiglet %v", err)
		log.Errorln(errStr)
	}
}

func RemoveDirContent(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := dir + "/" + file.Name()
		log.Infoln("RemoveDirConent found ", filename)
		err = os.RemoveAll(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run this:
//    DMDIR=/opt/zededa/bin/
//    ${DMDIR}/dnsmasq -b -C /var/run/zedrouter/dnsmasq.${BRIDGENAME}.conf
func startDnsmasq(bridgeName string) {

	log.Infof("startDnsmasq(%s)\n", bridgeName)
	hostDir := "--hostsdir=" + runDirname
	cfgPathname := dnsmasqConfigPath(bridgeName)
	name := "nohup"
	args := []string{
		"/opt/zededa/bin/dnsmasq",
		hostDir,
		"-C",
		cfgPathname,
	}
	log.Infof("Calling command %s %v\n", name, args)
	out, err := base.Exec(log, name, args...).CombinedOutput()
	if err != nil {
		log.Errorf("startDnsmasq: Failed starting dnsmasq for bridge %s (%s)",
			bridgeName, err)
	} else {
		log.Infof("startDnsmasq: Started dnsmasq with output: %s", out)
	}
}

//    pkill -u nobody -f dnsmasq.${BRIDGENAME}.conf
func stopDnsmasq(bridgeName string, printOnError bool, delConfiglet bool) {

	log.Infof("stopDnsmasq(%s)\n", bridgeName)
	pidfile := fmt.Sprintf("/var/run/dnsmasq.%s.pid", bridgeName)
	pidByte, err := ioutil.ReadFile(pidfile)
	if err != nil {
		log.Errorf("stopDnsmasq: pid file read error %v\n", err)
		return
	}
	pidStr := string(pidByte)
	pidStr = strings.TrimSuffix(pidStr, "\n")
	pidStr = strings.TrimSpace(pidStr)
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Errorf("stopDnsmasq: pid convert error %v\n", err)
		return
	}

	cfgFilename := dnsmasqConfigFile(bridgeName)
	pkillUserArgs("root", cfgFilename, printOnError)

	startCheckTime := time.Now()
	// check and wait until the process is gone or maximum of 60 seconds is reached
	for {
		p, err := os.FindProcess(pid)
		if err == nil {
			err = p.Signal(syscall.Signal(0))
			if err != nil {
				log.Infof("stopDnsmasq: kill process done for %d\n", pid)
				break
			} else {
				log.Infof("stopDnsmasq: wait for %d to finish\n", pid)
			}
			if time.Since(startCheckTime).Seconds() > 60 {
				log.Errorf("stopDnsmasq: kill dnsmasq on %s pid %d not finish in 60 seconds\n", bridgeName, pid)
				break
			}
			time.Sleep(1 * time.Second)
		} else {
			log.Infof("stopDnsmasq: find dnsmasq process %s error %v\n", pidStr, err)
			break
		}
	}

	err = os.Remove(pidfile)
	if err != nil && printOnError {
		log.Errorf("stopDnsmasq: remove pidfile %s error %v", pidfile, err)
	}

	if delConfiglet {
		deleteDnsmasqConfiglet(bridgeName)
	}
}

// checkAndPublishDhcpLeases needs to be called periodically since it
// refreshes the LastSeen and does garbage collection based on that timestamp
func checkAndPublishDhcpLeases(ctx *zedrouterContext) {
	changed := updateAllLeases(ctx)
	if !changed {
		return
	}
	// Walk all and update all which gained or lost a lease
	pub := ctx.pubAppNetworkStatus
	items := pub.GetAll()
	for _, st := range items {
		changed := false
		status := st.(types.AppNetworkStatus)
		for i := range status.UnderlayNetworkList {
			ulStatus := &status.UnderlayNetworkList[i]
			l := findLease(ctx, status.Key(), ulStatus.Mac, true)
			assigned := (l != nil)
			if ulStatus.Assigned != assigned {
				log.Infof("Changing(%s) %s mac %s to %t",
					status.Key(), status.DisplayName,
					ulStatus.Mac, assigned)
				ulStatus.Assigned = assigned
				changed = true
			}
		}
		if changed {
			publishAppNetworkStatus(ctx, &status)
		}
	}
}

// findLease returns a pointer so the caller can update the
// lease information
// Ignores a lease which has expired if ignoreExpired is set
func findLease(ctx *zedrouterContext, hostname string, mac string, ignoreExpired bool) *dnsmasqLease {
	for i := range ctx.dhcpLeases {
		l := &ctx.dhcpLeases[i]
		if l.Hostname != hostname {
			continue
		}
		if l.MacAddr != mac {
			continue
		}
		if ignoreExpired && l.LeaseTime.Before(time.Now()) {
			log.Warnf("Ignoring expired lease: %v", *l)
			return nil
		}
		log.Debugf("Found %v", *l)
		return l
	}
	log.Debugf("Not found %s/%s", hostname, mac)
	return nil
}

// addOrUpdateLease returns true if something changed
func addOrUpdateLease(ctx *zedrouterContext, lease dnsmasqLease) bool {
	l := findLease(ctx, lease.Hostname, lease.MacAddr, false)
	if l == nil {
		ctx.dhcpLeases = append(ctx.dhcpLeases, lease)
		log.Infof("Adding lease %v", lease)
		return true
	} else if !cmp.Equal(*l, lease) {
		log.Infof("Updating lease %v with %v", *l, lease)
		*l = lease
		return true
	} else {
		return false
	}
}

// markRemoveLease will fatal if the lease does not exist
// Merely marks for removal; see purgeRemovedLeases
func markRemoveLease(ctx *zedrouterContext, hostname string, mac string) {
	l := findLease(ctx, hostname, mac, false)
	if l == nil {
		log.Fatalf("Lease not found %s/%s", hostname, mac)
	}
	log.Infof("Removing lease %v", l)
	l.Remove = true
}

// purgeRemovedLeases does the actual removal of the marked entries
func purgeRemovedLeases(ctx *zedrouterContext) {
	var newLeases []dnsmasqLease
	var removed = 0
	for _, lease := range ctx.dhcpLeases {
		if lease.Remove {
			removed++
			continue
		}
		newLeases = append(newLeases, lease)
	}
	ctx.dhcpLeases = newLeases
	log.Infof("purgeRemovedLeases removed %d", removed)
	// XXX change to log.Debugf
	log.Infof("XXX after purgeRemovedLeases %v", ctx.dhcpLeases)
}

type dnsmasqLease struct {
	BridgeName string
	LastSeen   time.Time // For garbage collection
	Remove     bool      // Marked for removal

	// From leases file
	LeaseTime time.Time
	MacAddr   string
	IPAddr    string
	Hostname  string
}

const leaseGCTime = 5 * time.Minute

// updateAllLeases maintains ctx.dhcpLeases. It reads all of the leases and
// adds new ones immediately. It removes an old old if it haven't seen in
// 30 seconds (this is to handle file truncation/rewrite by dnsmasq) This
// timer assumes the function is called very 10 seconds.
// Returns true if there was a change to at least one lease
func updateAllLeases(ctx *zedrouterContext) bool {
	changed := false
	bridgeNames := dnsmasqBridgeNames()
	log.Debugf("bridgeNames: %v", bridgeNames)
	for _, bridgeName := range bridgeNames {
		leases, err := readLeases(bridgeName)
		if err != nil {
			log.Warnf("readLeases(%s) failed: %s", bridgeName, err)
			continue
		}
		log.Debugf("read leases(%s) %v", bridgeName, leases)
		// Add any new ones
		for _, l := range leases {
			if addOrUpdateLease(ctx, l) {
				changed = true
			}
		}
	}
	// Look for any old leases
	// XXX time limit is 5 minutes and we are assuming we get called
	// every 10 seconds or so.
	removed := false
	for _, l := range ctx.dhcpLeases {
		if time.Since(l.LastSeen) <= leaseGCTime ||
			time.Since(l.LeaseTime) <= leaseGCTime {
			continue
		}
		log.Infof("lease %v garbage collected: lastSeen %v ago, lease expiry %v ago",
			l, time.Since(l.LastSeen), time.Since(l.LeaseTime))
		markRemoveLease(ctx, l.Hostname, l.MacAddr)
		changed = true
		removed = true
	}
	if removed {
		purgeRemovedLeases(ctx)
	}
	return changed
}

// readLeases returns a slice of structs with mac, IP, uuid
// XXX Do we need to handle file which is deleted when bridge/networkinstance is deleted?
//
// Example content of leasesFile
// 1560664900 00:16:3e:00:01:01 10.1.0.3 63120af3-42c4-4d84-9faf-de0582d496c2 *
func readLeases(bridgeName string) ([]dnsmasqLease, error) {

	var leases []dnsmasqLease
	leasesFile := dnsmasqLeasePath(bridgeName)
	info, err := os.Stat(leasesFile)
	if err != nil {
		return leases, err
	}
	fileDesc, err := os.Open(leasesFile)
	if err != nil {
		return leases, err
	}
	reader := bufio.NewReader(fileDesc)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Errorln("ReadString ", err)
				return leases, err
			}
			break
		}
		// remove trailing "/n" from line
		line = line[0 : len(line)-1]

		// Should have 5 space-separated fields. We only use 4.
		tokens := strings.Split(line, " ")
		if len(tokens) < 4 {
			log.Errorf("Less than 4 fields in leases file: %v",
				tokens)
			continue
		}
		i, err := strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			log.Errorf("Bad unix time %s: %s", tokens[0], err)
			i = 0
		}
		lease := dnsmasqLease{
			BridgeName: bridgeName,
			LastSeen:   info.ModTime(),
			LeaseTime:  time.Unix(i, 0),
			MacAddr:    tokens[1],
			IPAddr:     tokens[2],
			Hostname:   tokens[3],
		}
		leases = append(leases, lease)
	}
	return leases, nil
}

// When we restart dnsmasq with smaller changes like chaging DNS server
// file configuration, we should not delete the hosts configuration for
// that has the IP address allotment information.
func deleteOnlyDnsmasqConfiglet(bridgeName string) {

	log.Infof("deleteOnlyDnsmasqConfiglet(%s)\n", bridgeName)
	cfgPathname := dnsmasqConfigPath(bridgeName)
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			errStr := fmt.Sprintf("deleteOnlyDnsmasqConfiglet %v",
				err)
			log.Errorln(errStr)
		}
	}
}
