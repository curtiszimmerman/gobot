/**
 * @project GoBot
 * GoBot IRC channel bot written in Go.
 * @file gobot.go
 * primary application driver
 * @author curtis zimmerman
 * @contact hey@curtisz.com
 * @license MIT
 * @version 0.0.1a
 */

package main

import (
	//"flag"
	. "fmt"
	"net"
	"os"
	"strconv"
)

type Connection struct {
	host   *net.IPAddr
	host_s string
	port   float64
	port_s string
}

func connect(cx Connection) Connection {
	return cx
}

func options() Connection {
	//flag.StringVar(&host, "host", "irc.freenode.net", "remote IRC server (default irc.freenod.net)")
	//flag.IntVar(&port, "port", 6667, "remote IRC port (default 6697)")
	//flag.Parse()
	if len(os.Args) != 3 {
		usage()
	}
	host_s, port_s := os.Args[1], os.Args[2]
	port, err := strconv.ParseFloat(port_s, 64)
	if err != nil {
		Printf("[!] Cannot parse port: ", err.Error(), " Dying...\n")
		os.Exit(1)
	}
	if &port == nil {
		Printf("[!] Invalid port! Dying...\n")
		os.Exit(1)
	}
	host, err := net.ResolveIPAddr("ip", host_s)
	if &host == nil {
		Printf("[!] Could not resolve address! Dying...\n")
		os.Exit(1)
	}
	Printf("[+] Application initialized...\n")
	cx := Connection{host: host, host_s: host_s, port: port, port_s: port_s}
	return cx
}

func usage() {
	Printf("IRC bot written in Go by curtisz\n")
	Printf("(https://github.com/curtiszimmerman/gobot)\n")
	Printf("Released under MIT license (C) 2014\n")
	Printf("\nUsage: %s [OPTION]... HOST [PORT]\n", os.Args[0])
	Printf("  -l logfile		log to specified file (not yet implemented)\n\n")
	os.Exit(1)
}

func version() {
	const script_name, version_pattern, version_release string = "GoBot", "%v v%v.%v.%v%v\n", "a"
	const version_major, version_minor, version_build uint = 0, 1, 0
	Printf(version_pattern, script_name, version_major, version_minor, version_build, version_release)
}

func main() {
	version()
	cx := options()
	// connect
	Printf("[+] Connecting to %v:%v (%v:%v)...\n", cx.host_s, cx.port_s, cx.host.String(), cx.port)
	connect(cx)
}
