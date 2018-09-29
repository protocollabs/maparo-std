# TCP Goodput

## Description

TCP Goodput can be considered as maparos iperf
replacement. It sent and receives TCP data. Beside
simple data exchange it has some novel features.

## Features

- Payload pattern: zeroized, random ASCII, full rand
- Configurable DSCP value (if OS support)
- Flexible traffic exchange configuration possibilies
- Setting No Delay
- Setting Maximum Setnet Size
- Goodput Limit configuration support. For TCP this is somewhat
  hacky. Especially for high data rates the userspace interaction can
  limit the overall system performance. So consider this feature as
  experimental
- Zerocopy mode uses snedfile (if OS support). Will not work in all
  other configuration options
- IPv6 flowlabel support
- Selectable congestion control algorithm
- Parallel Workers (thread support)


## Measurement Start Request


```
{
	# per default one TCP transmitter is started, to spawn exactly
	# to much threads are cores are available use "cores".
	# Use "threads" if you want to fully utilize all virtual cores,
	# including hyperthreads.
	# If the system has several sockets, all sockets are utilized for
	# "cores" and "threads".
	"streams" : "1"

	# payload pattern. Default is zeroized because we want to fullfill
	# the pipe and offload as much as possible. 
	"payload-pattern" : "zeroized"
}
```

## Measurement Start Reply

```
{
  "streams" :
  [
		{ "listen-port" : "<port>" }
	],
}
```


## Measurement Info Reply

```
{
  "streams" :
  [
	  {
		"timestamp-first" : "<maparo-time>"
		"timestamp-last"  : "<maparo-time>"
		"received-bytes" :      "<uint64_t>"
		}
	],
}
```

## Not Supported

- Ignore <n> seconds from start of measurement. This must be done by analysis tooling


