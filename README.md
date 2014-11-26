# gobot

GoBot is a simple IRC logging and channel administration bot written in Go.

[ ![Codeship Status for curtiszimmerman/gobot](https://codeship.com/projects/939b0170-5783-0132-1105-0e0cfcc5dfb4/status)](https://codeship.com/projects/49830)

### usage

```sh
./gobot [server] [port] [password] [nickname] [realname]
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
say | [channel] [message] | Force GoBot to send [message] to [channel]
die | [password] | Force Gobot to quit (requires [password])

### build from source

```sh
git clone https://github.com/curtiszimmerman/gobot /tmp/gobot
cd /tmp/gobot && go build
```

### contributors

Contributions to GoBot are welcome. However, changes which are mostly stylistic 
or semantic will not be accepted. If you submit a patch or pull request, please 
understand that accepting the changes may take time, at least until science 
invents 96-hour days.

### license

GoBot is (C) 2014 curtis zimmerman and released under GPLv3.