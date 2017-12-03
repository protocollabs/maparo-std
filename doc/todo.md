# Command Line Interface

## Syntax

Work in Progress

```
maparo <"module" | "campaign"> <name> <"client" | "server"> [<options>]
```

## Examples

```
maparo mod udp-rtt server config.port=8888
maparo mod udp-rtt client config.dst=1.1.1.1 config.port=8888
```

Print config values:

```
maparo mod udp-rtt client config-dump
```


Read config values from file:

```
maparo mod udp-rtt client config-read=/tmp/udp-rtt-client.json
```



To print help message about udp-rtt and links
to further information about client and server configs

```
maparo mod udp-rtt help
```


To print all command line argument and examples for client mode

```
maparo mod udp-rtt client help
```


Human readable or machine readable output (default human)

```
maparo mod udp-rtt client format=json config.dst=9.9.9.9
maparo mod udp-rtt client format=human config.dst=9.9.9.9
```


```
maparo mod udp-rtt server config.port=8888
^^^^^^^^^^^^^^^^^^^^^^^^^
```
- This has a pre-defined config
- mods comes as pairs, for client and server, always


maparo mod-config udp-rtt server dst=10

Read config, overwrite defaults, read args overwrite existing args - in this
order.

```
maparo mod-udp-ping client --verbose config=config-udp-ping.json addr=1.1.1.1
```

To print udp-rtt server config as JSON file to STDOUT


maparo campaign rrul server
maparo campaign rrul client config.dst=::1

Campaign rrul client:

[ delay:0.0, mod:udp-rtt, config.dst=$DST ]
[ delay:1.0, mod:udp-goodput, config.dst=$DST ]




# Control Plane

Maparo without arguments open up a control channel.
There are more use cases where a control channel is
possible as without. Exceptions are multicast modules.



maparo remote --key secret

# remote means that the other peer must operate in
# strict remote mode. E.g. the operation mode is
# dictated from the client. A remote process one client
# after another. Remote servers already processing a client
# all other incoming connects are dropped.
#
# The secret is just optional and is now has the same
# security level as the SNMP community string. Though,
# it is used as the encryption and prevents script kiddies
# to use a remote server somehow.
maparo --remote <ip:port> --remote-secret <secret> mod-udp-rtt --mode server --port 8888





# Campaing Config

```
{
	"config" : {
		"addr" : "::1",
		"port" : "6666"
	},

	"mods" : [
		{
			"delay" : "rand(0.0, 1.0)",
			"module" : "udp-ping",
			"mode" : "client",
			"args" : [
				"addr=$addr",
				"port=$port"
			]
		},
		{
			"delay" : "0",
			"module" : "udp-ping",
			"mode" : "client",
			"args" : [
				"addr=$addr",
				"port=$port"
			]
		}
	]
}
```


# Command Line Interface Equivalent to iperf, netperf, ...

## Iperf

### One way TCP, maximum throughput

Iperf:

```
# client
iperf -c <address>
# server
iperf -s -D
```

Maparo

```
# client
# server
maparo daemon
```


### TCP from Server to Client (reverse), maximum throughput

Iperf:

```
# client
iperf -d -c <address>
# server
iperf -s -D
```

Maparo

```
# client
maparo mod-tcp-goodput-reverse config.addr=<address>
# server
maparo daemon
```


### Parallel Bidirectional TCP, maximum throughput

Iperf:

```
# client
iperf -d -c <address>
# server
iperf -s -D
```

Maparo

```
# client
maparo campaign-tcp-bidirectinal config.addr=<address>
# server
maparo daemon
```

### Parallel Bidirectional TCP, maximum throughput, sequential

Iperf:

```
# client
iperf -d -c <address>
# server
iperf -s -D
```

Maparo

```
# client
maparo campaign-tcp-bidirectinal-sequential config.addr=<address>
# server
maparo daemon
```

