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

## Not Supported

- Ingnore <n> seconds from start of measuremtn. This must be done by analysis tooling
