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
	"os"
	"strconv"
)

type Connection struct {
	host string
	port float64
}

func connect(cx Connection) Connection {
	return cx
}

func options() (host, port string) {
	//flag.StringVar(&host, "host", "irc.freenode.net", "remote IRC server (default irc.freenod.net)")
	//flag.IntVar(&port, "port", 6667, "remote IRC port (default 6697)")
	//flag.Parse()
	if len(os.Args) != 3 {
		usage()
	}
	host, port = os.Args[1], os.Args[2]
	return
}

func usage() {
	version()
	Printf("IRC bot written in Go by curtisz\n")
	Printf("(https://github.com/curtiszimmerman/gobot)\n")
	Printf("Released under MIT license (C) 2014\n")
	Printf("\nUsage: gobot [OPTION]... HOST [PORT]\n")
	Printf("  -l logfile		log to specified file (not yet implemented)\n\n")
	os.Exit(1)
}

func version() {
	const script_name, version_pattern, version_release string = "GoBot", "%v v%v.%v.%v%v\n", "a"
	const version_major, version_minor, version_build uint = 0, 1, 0
	Printf(version_pattern, script_name, version_major, version_minor, version_build, version_release)
}

func main() {
	host, port_s := options()
	// connect
	version()
	Printf("[+] Connecting to %v:%v ...\n", host, port_s)
	port, err := strconv.ParseFloat(port_s, 64)
	if err != nil {
		Printf("Error parsing options: %v\n", err)
	}
	connect(Connection{host: host, port: port})
}
