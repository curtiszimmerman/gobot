# gobot

GoBot is a simple IRC logging and channel administration bot written in Go.

### usage

```sh
./gobot [SERVER] [PORT] [PASSWORD] [NICKNAME] [REALNAME]
```

The default server is "irc.freenode.net" and the default port is 6667. The default 
password to make GoBot die is "gobot". The default nickname is "gobot" and the default 
realname is "gobot-(version)" where (version) is the current version, e.g. "v0.0.1a".

To send commands to GoBot via IRC:

```sh
/msg gobot command [parameter]...
```

### commands

GoBot will accept the following commands:

Command | Parameters | Function
--- | --- | ---
help | n/a | Displays GoBot help message
say | <channel> <message> | Force GoBot to send <message> to <channel>
die | <password> | Force Gobot to quit (requires <password>)

### build from source

```sh
git clone https://github.com/curtiszimmerman/gobot /tmp/gobot
cd /tmp/gobot && go build
```
### license

GoBot is (C) 2014 Curtis Zimmerman and released under GPLv3.

### contributors

Contributions to GoBot are welcome. However, stylistic or semantic changes will not be accepted.