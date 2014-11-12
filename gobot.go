/**
 * @project GoBot
 * GoBot IRC channel bot written in Go.
 * @file gobot.go
 * primary application driver
 * @author curtis zimmerman
 * @contact hey@curtisz.com
 * @license MIT
 * @version 0.0.2a
 */

/*START gobot.go*/

package main

import (
	//"flag"
	"bufio"
	. "fmt"
	"net"
	"os"
	"strconv"
)

type Config struct {
	host   *net.IPAddr
	host_s string
	port   int64
	port_s string
}

type Server struct {
	instance string
}

// send a message to a channel
func (server *Server) message(channel, data string) bool {
	//@debug1
	Printf("[-] instance ID (%v) sending to channel (%v) message: %v", server.instance, channel, data)
	return false
}

// flush the bot and refresh
func (server *Server) flush() bool {
	return false
}

func connect(cx Config) net.Conn {
	conn, err := net.Dial("tcp", cx.host.String()+":"+cx.host.String())
	if err != nil {
		Printf("[!] Could not initiate connection: %v\n", err)
	}
	return conn
}

func options() Config {
	//flag.StringVar(&host, "host", "irc.freenode.net", "remote IRC server (default irc.freenod.net)")
	//flag.IntVar(&port, "port", 6667, "remote IRC port (default 6697)")
	//flag.Parse()
	if len(os.Args) != 3 {
		usage()
	}
	host_s, port_s := os.Args[1], os.Args[2]
	port, err := strconv.ParseInt(port_s, 10, 64)
	if err != nil {
		Printf("[!] Cannot parse port: %v", err.Error())
		os.Exit(1)
	}
	if &port == nil {
		Printf("[!] Could not parse port!\n")
		os.Exit(1)
	}
	if port < 0 || port > 65535 {
		Printf("[!] Port must be between 1 and 65535!\n")
		os.Exit(1)
	}
	host, err := net.ResolveIPAddr("ip", host_s)
	if err != nil {
		Printf("[!] Could not resolve address: %v\n")
		os.Exit(1)
	}
	Printf("[+] Application initialized...\n")
	cx := Config{host: host, host_s: host_s, port: port, port_s: port_s}
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
	const version_major, version_minor, version_build uint = 0, 0, 2
	Printf(version_pattern, script_name, version_major, version_minor, version_build, version_release)
}

func main() {
	version()
	cx := options()
	// connect
	Printf("[+] Connecting to %v:%v (%v:%v)...\n", cx.host_s, cx.port_s, cx.host.String(), cx.port)
	conn := connect(cx)
	connBuffer := bufio.NewReader(conn)
	for {
		// parse input
		str, err := connBuffer.ReadString('\n')
		if len(str) > 0 {
			Printf(str)
		}
		if err != nil {
			Printf("[!] Error receiving on socket: %v\n")
			os.Exit(1)
		}
		Printf("..................... next line .....................")
	}
}

/*EOF gobot.go*/
