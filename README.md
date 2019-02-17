# Maparo

## Abstract

> *Esperanto for Atlas*

Maparo is a network performance measurement protocol specification. Beside
iperf, netperf and other tools it just defines the protocol
specification - not one particular implementation. Similar to HTTP/2 (RFC 7540) or any
other networking protocol specification.

Maparo differiantiate between the control protocol (e.g. probe the server or
start measurement) and the measurement protocol (e.g. tcp stream with sequence
numbers). The later specification is done in seperate module specifcation.

Maparo was designed to be flexible and extensible. Maparo differiantiate
between the control protocol and the measurement protocol.  The control
protocol is keep as simple as possible and provides a basis functionality like
service discovery, measurement start request and so on. It provides the
transport layer for the module specific measurement protocol. The actual workhorses are
implemented in so called "maparo modules", they implement the building blocks
for measurements. Some of them are mandatory, many are optional and it is also
possible to develop completely proprietary modules.

There is one reference implementation: mapago (implemented in go, thus the
name). A Python implementation is also available (but protocol support is based
on older version of maparo). But you can - and should - program your own
implementation in your language of choice. No matter if it is GUI or command
line interface tool.

**KEEP IN MIND:** maparo protocol is not finalized yet. We do our best not to
change the existing specification, but we cannot rule it out.

## Introdcuction

### Modules

Modules can be manadory or optional. Modules itself can require mandatory
featureset and provide an optional featureset.

#### Mandatory Modules

Mandatory module feature functionality MUST be able to implemented at any
platform. Operating system specific features MUST NOT be required in a
Mandatory Module.

#### Optional Modules

Blessed and officially released modules. Implementation can implement these
modules if they want. If Optional Modulesa are implemented they MUST follow the
specification.

#### Unofficial Modules

Possibility to implement propritary modules.

### Time Format

All internally measured and transferred timevalue should use a realtime clock
(`CLOCK_REALTIME`). `CLOCK_MONOTONIC` principle be used if a clean synchronization
between client and server can de done. This is principle be true for remote mode.
But because the time synchronization is a) not that accurate as required and b) not
possible at all if operated in a non-remote mode.

The only solution is to ignore this within maparo. If a high resolution timing
analysis between client and server is required the only solution is to use GPS/PTP
for the time of measurement.

`CLOCK_MONOTONIC_RAW` would be fine if we can build upon maparo internal time
synchronization mechanisms - but we can't.

#### Exchanged Time Format

If time is exchanged via JSON the format MUST be UTC. The time resultion should
be in nanoseconds:

```
2017-12-16T12:32:42.763987000
```

Implementations SHOULD check the number of digits of the fractions. If the number
is six then microseconds is used. If 9 digits it should be interpreted as nanoseconds.

With Python3:

```python
import datetime
dt = datetime.datetime.utcnow()
print(dt.strftime('%Y-%m-%dT%H:%M:%S.%f'))
```

For Golang:

```go
import "time"
import "fmt"
t = time.Now().UTC()
fmt.Println(t.Format("2006-01-02T15:04:05.000000000"))
```

and reverse

```go
ret, err := time.Parse("2006-01-02T15:04:05.000000000" , t)
if err != nil {
	// do what you want
}
```

### Payload Pattern

Maparo pre-defines several payload pattern to be
used in modules.

The pattern is true for one "chunk". One chunk is one pre allocated buffer and
is typically one UDP packet or one large TCP chunk. Chunks are reused and
pattern are not recalculated (thus identical). This is for performance aspects
because randomizing and touching chunks are CPU intensive and may lower the
network performance. I don't see any network measurement advantageous where
recalculating is required. Often optimizer and gzip for UDP work on a packet
level and for TCP the chunk size can be quite large so there is no real problem
with this limitation.

#### Zero

Just 0 for the complete payload

Name: `zero`

#### Random ASCII (letter)

Randomized string with `a-zA-Z0-9`. No Unicode

Name: `random-ascii`

#### Random 

Random is a pure random byte generator.  The generator tries to use the most
cryptographic random bytes from the underlying operating system (e.g.
`/dev/random` seed combined with AES)

Name: `random`




- [command-line-interface.md](command-line-interface.md)

## Control Protocol

- [control-protocol.md](control-protocol.md)

## Modules

- [TCP Goodput (tcp-goodput)](mod-tcp-goodput.md)
- [UDP RTT (udp-rtt)](mod-udp-rtt.md)
- [UDP Ping (udp-ping)](mod-udp-ping.md)
- [UDP Mcast Spray (udp-mcast-spray)](mod-udp-mcast-spray.md)
- [UDP Goodput (udp-goodput)](mod-udp-goodput.md)

> Common used modules are named shorter/snappy, rarely use modules named
longer

## Campaigns

**WIP:** Campaigns are not shared or/and standardized. Campaings consists
of modules. Modules are standardized and can be queried. A campaign can
get supported modules from a peer and can start a campaign if all required
modules are supported by a peer/server.

## Misc

- There is not output in between measurement. Why: because the control
  channel is not used. Only after measurement the control channel is used
	to exchange the measurement data back to the client. The client is free
	to print status information if the cpu load is not noticable affected.

