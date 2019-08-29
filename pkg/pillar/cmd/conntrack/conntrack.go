// Copyright (c) 2019 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package conntrack

import (
	"flag"
	"fmt"
	"github.com/eriknordmark/netlink"
	log "github.com/sirupsen/logrus"
	"net"
	"syscall"
)

func Run() {
	// XXX curpartPtr := flag.String("c", "", "Current partition")
	delFlow := flag.Bool("D", false, "Delete flow")
	delSrcIP := flag.String("s", "", "Delete flow with Srouce IP")
	delProto := flag.Int("p", 0, "Delete flow with protocol ID")
	delFamily := flag.String("f", "", "Delete flow with ipv6")
	delPort := flag.Int("P", 0, "Delete flow with port number")
	flag.Parse()

	// conntrack [-D <-s address> [-p proto]]
	if *delFlow {
		if *delSrcIP != "" {
			var proto uint8
			var src net.IP
			var port uint16
			var family netlink.InetFamily

			family = syscall.AF_INET
			src = net.ParseIP(*delSrcIP)
			if *delProto != 0 {
				proto = uint8(*delProto)
			}
			if *delFamily == "ipv6" {
				family = syscall.AF_INET6
			}
			if *delPort != 0 {
				port = uint16(*delPort)
			}

			number, err := netlink.ConntrackDeleteIPSrc(netlink.ConntrackTable, family, src, proto, port, true)
			if err != nil {
				log.Println("ConntrackDeleteIPSrc error:", err)
			} else {
				fmt.Printf("ConntrackDeleteIPSrc: deleted %d flow\n", number)
			}
			return
		}
		fmt.Println("Usage: Conntrack -D -s IP-Address [-p Protocol][-P port][-f ipv6]")
		return
	}
	// XXX args := flag.Args()
	// XXX curpart := *curpartPtr
	res, err := netlink.ConntrackTableList(netlink.ConntrackTable, syscall.AF_INET)
	if err != nil {
		log.Println("ContrackTableList", err)
	} else {
		for i, entry := range res {
			fmt.Printf("[%d]: %s\n", i, entry.String())
			fmt.Printf("[%d]: forward packets %d bytes %d\n", i,
				entry.Forward.Packets, entry.Forward.Bytes)
			fmt.Printf("[%d]: reverse packets %d bytes %d\n", i,
				entry.Reverse.Packets, entry.Reverse.Bytes)
		}
	}
	res, err = netlink.ConntrackTableList(netlink.ConntrackTable, syscall.AF_INET6)
	if err != nil {
		log.Println("ContrackTableList", err)
	} else {
		for i, entry := range res {
			fmt.Printf("[%d]: %s\n", i, entry.String())
			fmt.Printf("[%d]: forward packets %d bytes %d\n", i,
				entry.Forward.Packets, entry.Forward.Bytes)
			fmt.Printf("[%d]: reverse packets %d bytes %d\n", i,
				entry.Reverse.Packets, entry.Reverse.Bytes)
		}
	}
}
