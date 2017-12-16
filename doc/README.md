# Maparo

## Modules

- [TCP Goodput (tcp-goodput)](mod-tcp-goodput.md)
- [UDP RTT (udp-rtt)](mod-udp-rtt.md)
- [UDP Ping (udp-ping)](mod-udp-ping.md)
- [UDP Mcast Spray (udp-mcast-spray)](mod-udp-mcast-spray.md)

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

All internally measured and transferred timevalue should use a realtime clock
(`CLOCK_REALTIME`). `CLOCK_MONOTONIC` principle be used if a clean synchronization
between client and server can de done. This is principle be true for remote mode.
But because the time synchronization is a) not that accurate as required and b) not
possible at all if operated in a non-remote mode.

The only solution is to ignore this within maparo. If a high resolution timing
analysis between client and server is required the only solution is to use GPS
mouses and disable gpsd for the time of measurement to do not risk NTP
adjustments.

`CLOCK_MONOTONIC_RAW` would be fine if we can build upon maparo internal time
synchronization mechanisms - but we can't.

Use: `

### Exchanged Time Format

If time is exchanged via JSON the format must be UTC and with a resulution
of microseconds:

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


