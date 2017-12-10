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

### Exchanged Time Format

If time is exchanged via JSON the format must be UTC and with a resulution
of nanosecond:

```
2017-05-14T23:55:00.123456789Z
```

# Cross Platform Functionality

Not all features are supported at every platform. These features
must be described and named explicitly.

## Principles

- Keep the serve size simple and configuration free from configuration
  effort. The configuration should be keep as simple as possible at the
  server side, should be done on the client side and as automatically as
  possible.

