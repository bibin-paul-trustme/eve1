// Copyright (c) 2017-2018 Zededa, Inc.
// All rights reserved.

// dnsmasq configlets for overlay and underlay interfaces towards domU

package zedrouter

import (
	"fmt"
	"github.com/zededa/go-provision/agentlog"
	"github.com/zededa/go-provision/types"
	"log"
	"net"
	"os"
	"os/exec"
)

const dnsmasqStatic = `
# Automatically generated by zedrouter
except-interface=lo
bind-interfaces
log-queries
log-dhcp
no-hosts
no-ping
bogus-priv
stop-dns-rebind
rebind-localhost-ok
neg-ttl=10
`

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
	dhcphostsDir := globalRunDirname + "/dhcp-hosts." + bridgeName
	return dhcphostsDir
}

// When we create a linux bridge we set this up
// Also called when we need to update the ipsets
func createDnsmasqConfiglet(bridgeName string, bridgeIPAddr string,
	netconf *types.NetworkObjectConfig, hostsDir string,
	ipsets []string) {

	if debug {
		log.Printf("createDnsmasqConfiglet: %s netconf %v\n",
			bridgeName, netconf)
	}
	cfgPathname := dnsmasqConfigPath(bridgeName)
	// Delete if it exists
	if _, err := os.Stat(cfgPathname); err == nil {
		if err := os.Remove(cfgPathname); err != nil {
			log.Println(err)
		}
	}
	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("os.Create for ", cfgPathname, err)
	}
	defer file.Close()

	// Create a dhcp-hosts directory to be used when hosts are added
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	ensureDir(dhcphostsDir)

	file.WriteString(dnsmasqStatic)
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
	// static router defined in the NetworkObjectConfig).
	// To support airgap networks we interpret gateway=0.0.0.0
	// to not advertize ourselves as a router. Also,
	// if there is not an explicit dns server we skip
	// advertising that as well.
	advertizeRouter := true
	var router string
	if netconf.Gateway != nil {
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
	advertizedDns := false
	for _, ns := range netconf.DnsServers {
		advertizedDns = true
		file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server,%s\n",
			ns.String()))
	}
	if netconf.NtpServer != nil {
		file.WriteString(fmt.Sprintf("dhcp-option=option:ntp-server,%s\n",
			netconf.NtpServer.String()))
	}
	if netconf.Subnet.IP != nil {
		ipv4Netmask = net.IP(netconf.Subnet.Mask).String()
		file.WriteString(fmt.Sprintf("dhcp-option=option:netmask,%s\n",
			ipv4Netmask))
	}
	if advertizeRouter {
		// IPv6 XXX needs to be handled in radvd
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router,%s\n",
				router))
		}
	} else {
		log.Printf("createDnsmasqConfiglet: no router\n")
		if !isIPv6 {
			file.WriteString(fmt.Sprintf("dhcp-option=option:router\n"))
		}
		if !advertizedDns {
			// Handle isolated network by making sure
			// we are not a DNS server. Can be overridden
			// with the DnsServers above
			log.Printf("createDnsmasqConfiglet: no DNS server\n")
			file.WriteString(fmt.Sprintf("dhcp-option=option:dns-server\n"))
		}
	}
	if netconf.DhcpRange.Start != nil {
		dhcpRange = netconf.DhcpRange.Start.String()
	}
	if isIPv6 {
		file.WriteString(fmt.Sprintf("dhcp-range=::,static,0,10m\n"))
	} else {
		file.WriteString(fmt.Sprintf("dhcp-range=%s,static,%s,10m\n",
			dhcpRange, ipv4Netmask))
	}
}

func addhostDnsmasq(bridgeName string, appMac string, appIPAddr string,
	hostname string) {

	log.Printf("addhostDnsmasq(%s, %s, %s, %s)\n", bridgeName, appMac,
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
		log.Fatal(err)
	}
	defer file.Close()
	if isIPv6 {
		file.WriteString(fmt.Sprintf("%s,[%s],%s\n",
			appMac, appIPAddr, hostname))
	} else {
		file.WriteString(fmt.Sprintf("%s,id:*,%s,%s\n",
			appMac, appIPAddr, hostname))
	}
}

func removehostDnsmasq(bridgeName string, appMac string, appIPAddr string) {
	log.Printf("removehostDnsmasq(%s, %s, %s)\n",
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
		log.Printf("removehostDnsmasq(%s, %s) failed: %s\n",
			bridgeName, appMac, err)
		return
	}
	if err := os.Remove(cfgPathname); err != nil {
		log.Println(err)
	}
}

func deleteDnsmasqConfiglet(bridgeName string) {
	if debug {
		log.Printf("deleteDnsmasqConfiglet(%s)\n", bridgeName)
	}
	cfgPathname := dnsmasqConfigPath(bridgeName)
	if err := os.Remove(cfgPathname); err != nil {
		log.Println(err)
	}
	dhcphostsDir := dnsmasqDhcpHostDir(bridgeName)
	if err := os.RemoveAll(dhcphostsDir); err != nil {
		log.Println(err)
	}
	// XXX also delete hostsDir?
}

// Run this:
//    DMDIR=/opt/zededa/bin/
//    ${DMDIR}/dnsmasq -b -C /var/run/zedrouter/dnsmasq.${BRIDGENAME}.conf
func startDnsmasq(bridgeName string) {
	if debug {
		log.Printf("startDnsmasq(%s)\n", bridgeName)
	}
	cfgPathname := dnsmasqConfigPath(bridgeName)
	name := "nohup"
	//    XXX currently running as root with -d above
	args := []string{
		"/opt/zededa/bin/dnsmasq",
		"-d",
		"-C",
		cfgPathname,
	}
	logFilename := fmt.Sprintf("dnsmasq.%s", bridgeName)
	logf, err := agentlog.InitChild(logFilename)
	if err != nil {
		log.Fatalf("startDnsmasq agentlog failed: %s\n", err)
	}
	cmd := exec.Command(name, args...)
	cmd.Stderr = logf
	if debug {
		log.Printf("Calling command %s %v\n", name, args)
	}
	go cmd.Run()
}

//    pkill -u nobody -f dnsmasq.${BRIDGENAME}.conf
func stopDnsmasq(bridgeName string, printOnError bool) {
	if debug {
		log.Printf("stopDnsmasq(%s)\n", bridgeName)
	}
	cfgFilename := dnsmasqConfigFile(bridgeName)
	// XXX currently running as root with -d above
	pkillUserArgs("root", cfgFilename, printOnError)
}
