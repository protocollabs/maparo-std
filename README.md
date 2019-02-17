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

### Time

- [Time](time.md)


# Cross Platform Functionality

Not all features are supported at every platform. These features
must be described and named explicitly.

## Principles

- Keep the serve size simple and configuration free from configuration
  effort. The configuration should be keep as simple as possible at the
  server side, should be done on the client side and as automatically as
  possible.


# Shared Functionality

## Payload Pattern

Maparo pre-defines several payload pattern to be
used in each module.

The pattern is true for one "chunk". One chunk is one pre allocated buffer and
is typically one UDP packet or one large (max 4 GB) TCP chunk. Chunks are
reused and pattern are not recalculated -> identical. This is for performance
aspects because randomizing and touching chunks are CPU intensive and may lower
the network performance. I don't see any network measurement advantageous where
recalculating is required. Often optimizer and gzip for UDP work on a packet
level and for TCP the chunk size can be quite large so there is no real problem
with this limitation.

# Zero

Just 0 for the complete payload

Name: `zero`

### Random ASCII (letter)

Randomized string with `a-zA-Z0-9`. No Unicode

Name: `random-ascii`

### Random 

Random is a pure random byte.

Name: `random`

The generator tries to use the most cryptographic random bytes from the
underlying operating system (e.g. `/dev/random` seed combined with AES)


