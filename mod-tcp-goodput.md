# TCP Goodput

## Description

TCP Goodput can be considered as maparos iperf replacement. It sent and
receives TCP data. Beside simple data exchange it has some novel features.

> Note: `tcp-goodput` is limited to receive data send from client toward
> server. The server simple discards the received data and never sends data
> back. For additional measurements new modules must be speciefied and added.
> This is an rather simple module. Implementable with reduced effort with all
> major programming languages and operating systems.

## Features

- Payload pattern: zeroized, random ASCII, full rand (client side)
- Configurable DSCP value (if OS support, client side)
- Flexible traffic exchange configuration possibilies
- Setting No Delay
- Setting Maximum Segmet Size
- Goodput Limit configuration support. For TCP this is somewhat
  hacky. Especially for high data rates the userspace interaction can
  limit the overall system performance. So consider this feature as
  experimental
- Zerocopy mode uses snedfile (if OS support). Will not work in all
  other configuration options
- IPv6 flowlabel support
- Selectable congestion control algorithm
- Parallel Workers (thread support)

## Name

This module is standardized with the name:

```
tcp-goodput
```

## Measurement Start Request

### Info Reply

`tcp-goodput` has no additional information for the client. The `tcp-goodput`
dictionary MUST be empty.

E.g.

```
[]
  "id" : "hostname=uuid",
  "seq-rp" : <uint64_t>
  "modules" : {
     "tcp-goodput" : { "cores" : "4" },
  }
[]
```

### Measurement Start Request

```
{
	"streams" : "1"
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


