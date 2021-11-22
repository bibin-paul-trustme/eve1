// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package nim

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	dns "github.com/Focinfi/go-dns-resolver"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
)

const (
	minTTLSec       int = 30
	maxTTLSec       int = 3600
	extraSec        int = 10
	etcHostFileName     = "/etc/hosts"
	tmpHostFileName     = "/tmp/etchosts"
	resolvFileName      = "/etc/resolv.conf"
)

// go routine for dns query to the controller
func queryControllerDNS(ps *pubsub.PubSub) {
	var etchosts, controllerServer []byte
	var ttlSec int
	var ipaddrCached string

	if _, err := os.Stat(etcHostFileName); err == nil {
		etchosts, err = ioutil.ReadFile(etcHostFileName)
		if err == nil {
			controllerServer, _ = ioutil.ReadFile(types.ServerFileName)
			controllerServer = bytes.TrimSuffix(controllerServer, []byte("\n"))
			if bytes.Contains(controllerServer, []byte(":")) {
				serverport := bytes.Split(controllerServer, []byte(":"))
				if len(serverport) == 2 {
					controllerServer = serverport[0]
				}
			}
		}
	}

	if len(controllerServer) == 0 {
		log.Errorf("can't read /etc/hosts or server file")
		return
	}

	dnsTimer := time.NewTimer(time.Duration(minTTLSec) * time.Second)

	wdName := agentName + "dnsQuery"
	stillRunning := time.NewTicker(stillRunTime)
	ps.StillRunning(wdName, warningTime, errorTime)
	ps.RegisterFileWatchdog(wdName)

	for {
		select {
		case <-dnsTimer.C:
			// base on ttl from server dns update frequency for controller IP resolve
			// even if the dns server implementation returns the remaining value of the TTL it caches,
			// it will still work.
			ipaddrCached, ttlSec = controllerDNSCache(etchosts, controllerServer, ipaddrCached)
			dnsTimer = time.NewTimer(time.Duration(ttlSec) * time.Second)

		case <-stillRunning.C:
		}
		ps.StillRunning(wdName, warningTime, errorTime)
	}
}

// periodical cache the controller DNS resolution into /etc/hosts file
// it returns the cached ip string, and TTL setting from the server
func controllerDNSCache(etchosts, controllerServer []byte, ipaddrCached string) (string, int) {
	if len(etchosts) == 0 || len(controllerServer) == 0 {
		return ipaddrCached, maxTTLSec
	}

	// Check to see if the server domain is already in the /etc/hosts as in eden,
	// then skip this DNS queries
	if ipaddrCached == "" {
		hostsEntries := bytes.Split(etchosts, []byte("\n"))
		for _, entry := range hostsEntries {
			fields := bytes.Fields(entry)
			if len(fields) == 2 {
				if bytes.Compare(fields[1], controllerServer) == 0 {
					log.Tracef("server entry %s already in /etc/hosts, skip", controllerServer)
					return ipaddrCached, maxTTLSec
				}
			}
		}
	}

	var nameServers []string
	dnsServer, _ := ioutil.ReadFile(resolvFileName)
	dnsRes := bytes.Split(dnsServer, []byte("\n"))
	for _, d := range dnsRes {
		d1 := bytes.Split(d, []byte("nameserver "))
		if len(d1) == 2 {
			nameServers = append(nameServers, string(d1[1]))
		}
	}
	if len(nameServers) == 0 {
		nameServers = append(nameServers, "8.8.8.8")
	}

	if _, err := os.Stat(tmpHostFileName); err == nil {
		os.Remove(tmpHostFileName)
	}

	var newhosts []byte
	var gotipentry bool
	var lookupIPaddr string
	var ttlSec int

	domains := []string{string(controllerServer)}
	dtypes := []dns.QueryType{dns.TypeA}
	for _, nameServer := range nameServers {
		resolver := dns.NewResolver(nameServer)
		resolver.Targets(domains...).Types(dtypes...)

		res := resolver.Lookup()
		for target := range res.ResMap {
			for _, r := range res.ResMap[target] {
				dIP := net.ParseIP(r.Content)
				if dIP == nil {
					continue
				}
				lookupIPaddr = dIP.String()
				ttlSec = getTTL(r.Ttl)
				if ipaddrCached == lookupIPaddr {
					log.Tracef("same IP address %s, return", lookupIPaddr)
					return ipaddrCached, ttlSec
				}
				serverEntry := fmt.Sprintf("%s %s\n", lookupIPaddr, controllerServer)
				newhosts = append(etchosts, []byte(serverEntry)...)
				gotipentry = true
				// a rare event for dns address change, log it
				log.Noticef("dnsServer %s, ttl %d, entry add to /etc/hosts: %s", nameServer, ttlSec, serverEntry)
				break
			}
			if gotipentry {
				break
			}
		}
		if gotipentry {
			break
		}
	}

	if ipaddrCached == lookupIPaddr {
		return ipaddrCached, minTTLSec
	}
	if !gotipentry { // put original /etc/hosts file back
		newhosts = append(newhosts, etchosts...)
	}

	ipaddrCached = ""
	err := ioutil.WriteFile(tmpHostFileName, newhosts, 0644)
	if err == nil {
		if err := os.Rename(tmpHostFileName, etcHostFileName); err != nil {
			log.Errorf("can not rename /etc/hosts file %v", err)
		} else {
			if gotipentry {
				ipaddrCached = lookupIPaddr
			}
			log.Tracef("append controller IP %s to /etc/hosts", lookupIPaddr)
		}
	} else {
		log.Errorf("can not write /tmp/etchosts file %v", err)
	}

	return ipaddrCached, ttlSec
}

func getTTL(ttl time.Duration) int {
	ttlSec := int(ttl.Seconds())
	if ttlSec < minTTLSec {
		// this can happen often, when the dns server returns ttl being the remaining value
		// of it's own cached ttl, we set it to minTTLSec and retry. Next time will get the
		// upper range value of it's remaining ttl.
		ttlSec = minTTLSec
	} else if ttlSec > maxTTLSec {
		ttlSec = maxTTLSec
	}

	// some dns server returns actual remaining time of TTL, to avoid next time
	// get 0 or 1 those numbers, add some extra seconds
	return ttlSec + extraSec
}
