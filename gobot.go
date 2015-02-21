/**
 * @project GoBot
 * GoBot IRC channel bot written in Go.
 * @file gobot.go
 * primary application driver
 * @author curtis zimmerman
 * @contact hey@curtisz.com
 * @license AGPL
 * @version 0.0.2a
 */

/*START gobot.go*/

package main

import (
	//"flag"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"strconv"
	"strings"
)

/*\
|*| variables
\*/
var (
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var (
	version_major, version_minor, version_buildd, version_phase = 0, 0, 2, "a"
)

type Addresses struct {
	host   *net.IPAddr
	host_s string
	port   int64
	port_s string
}

type AuthServices struct {
	enabled bool
	broker  string
	pass    string
	user    string
}

type Client struct {
	addresses *Addresses
	inbound   chan string
	outbound  chan string
	reader    *bufio.Reader
	writer    *bufio.Writer
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err != nil {
			Warning.Printf("Could not read input: %v", err)
		}
		client.inbound <- line
	}
}

func (client *Client) Write() {
	for data := range client.outbound {
		client.writer.WriteString(data)
		client.writer.Flush()
	}
}

func GetClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	client := &Client{
		inbound:  make(chan string),
		outbound: make(chan string),
		reader:   reader,
		writer:   writer}
	client.Listen()
	return client
}

// this excellent pattern comes from: www.goinggo.net/2013/11/using-log-package-in-go.html
func Init(traceHandle io.Writer, debugHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(debugHandle, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "[INFO ] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "[WARNG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

type Input struct {
	prefix  string
	command string
	params  string
}

type Server struct {
	instance string
}

// send a message to a channel
func (server *Server) message(channel, data string) bool {
	//@debug1
	fmt.Printf("[-] instance ID (%v) sending to channel (%v) message: %v", server.instance, channel, data)
	return false
}

// flush the bot and refresh
func (server *Server) flush() bool {
	return false
}

type Settings struct {
	altnick    string
	altaltnick string
	authpass   string
	authuser   string
	channel    string
	nickname   string
	username   string
	password   string
	realname   string
	version    string
}

type Version struct {
	major int
	minor int
	build int
	phase string
}

/*\
|*| functions
\*/
func connect(addr *Addresses) net.Conn {
	conn, err := net.Dial("tcp", addr.host_s+":"+addr.port_s)
	if err != nil {
		fmt.Printf("[!] Could not initiate connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func options(version *Version) (*Addresses, *Settings) {

	var config, host_s, port_s, nick, channel, password string
	var auth bool

	flag.StringVar(&config, "config", "", "configuration file to use.")
	flag.StringVar(&host_s, "host", "irc.freenode.net", "remote IRC server.")
	flag.StringVar(&port_s, "port", "6667", "remote IRC port.")
	flag.StringVar(&nick, "nick", "gobot", "nickname to use.")
	flag.StringVar(&channel, "channel", "gobot-test", "channel to connect to.")
	flag.StringVar(&password, "password", "", "server passowrd to connect with.")
	flag.BoolVar(&auth, "auth", false, "auth to services. boolean.")
	flag.Parse()

	/*if host_s == "" || port_s == "" {
		usage()
	}
	if len(os.Args) > 4 {
		host_s, port_s, nick, channel = os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	}
	if len(os.Args) > 5 {
		password = os.Args[5]
	}*/
	port, err := strconv.ParseInt(port_s, 10, 64)
	if err != nil {
		Warning.Printf("could not parse port: %v\n", err)
	}
	if &port == nil {
		Error.Printf("could not parse port: %v\n", err)
		os.Exit(1)
	}
	if port < 0 || port > 65535 {
		Error.Printf("port must be between 1 and 65535\n")
		os.Exit(1)
	}
	host, err := net.ResolveIPAddr("ip", host_s)
	if err != nil {
		Warning.Printf("could not resolve address: %v\n")
		os.Exit(1)
	}
	if &nick == nil {
		Info.Printf("could not parse nickname: %v\n")
		nick = "gobot"
	}
	if &password == nil {
		Info.Printf("not using password...\n")
	}
	/*authsvc := &AuthServices{
		enabled: auth,
	}*/
	v := strconv.Itoa(version.major) + "." + strconv.Itoa(version.minor) + "." + strconv.Itoa(version.build) + version.phase
	settings := &Settings{
		//auth:       auth,
		nickname:   nick,
		altnick:    nick + "_",
		altaltnick: nick + "__",
		channel:    channel,
		password:   password,
		realname:   nick + v,
		username:   nick,
		version:    v}
	addr := &Addresses{host: host, host_s: host_s, port: port, port_s: port_s}
	Info.Printf("application initialized...\n")
	return addr, settings
}

func usage() {
	fmt.Printf("IRC bot written in Go by curtisz\n")
	fmt.Printf("(https://github.com/curtiszimmerman/gobot)\n")
	fmt.Printf("Released under MIT license (C) 2014-2015\n")
	fmt.Printf("\nUsage: %s [OPTION]... HOST [PORT] [NICK] [CHANNEL] [PASSWORD]\n", os.Args[0])
	fmt.Printf("  -l logfile		log to specified file (not yet implemented)\n\n")
	os.Exit(1)
}

func version() *Version {
	const script_name, version_pattern, version_phase string = "GoBot", "%v v%v.%v.%v%v\n", "a"
	const version_major, version_minor, version_build int = 0, 0, 2
	fmt.Printf(version_pattern, script_name, version_major, version_minor, version_build, version_phase)
	version := &Version{
		major: version_major,
		minor: version_minor,
		build: version_build,
		phase: version_phase}
	return version
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	var logfile string
	flag.StringVar(&logfile, "l", "", "log output to specified file (not yet implemented)")

	version := version()
	addr, settings := options(version)
	// connect
	Info.Printf("connecting to %v:%v (%v:%v)...\n", addr.host_s, addr.port_s, addr.host.String(), addr.port)
	conn := connect(addr)
	client := GetClient(conn)
	client.addresses = addr
	inbound := <-client.inbound

	time.Sleep(2)
	if settings.password != "" {
		Info.Printf("sending server password info [hidden]")
		pass := "PASS " + settings.password + "\n"
		client.outbound <- pass
	}
	nick := "NICK " + settings.nickname + "\n"
	Info.Printf("sending nickname info [%v]", nick)
	client.outbound <- nick
	// USER gobot gobot irc.freenode.net : gobotreal
	ident := "USER " + settings.nickname + " " + settings.nickname + " " + addr.host_s + " :" + settings.realname + "\n"
	Info.Printf("sending ident info [%v]", ident)
	client.outbound <- ident

	for {
		// parse input
		inbound = <-client.inbound
		if len(inbound) > 0 {
			Info.Printf(inbound)
		}
		message := strings.SplitN(inbound, ":", 3)
		msg := Input{prefix: "", command: "", params: ""}
		if len(message) > 0 {
			msg.prefix = message[0]
		}
		if len(message) > 1 {
			msg.command = message[1]
		}
		if len(message) > 2 {
			msg.params = message[2]
		}
		if msg.prefix == "" {

		}
		/*
			if strings.Count(msg.command, ":") {
				// handle PING by sending PONG :hostname.example.com
				response = "PONG :"
			}
		*/
		fmt.Printf("\n..................... next line .....................\n")
	}
}

/*EOF gobot.go*/
