# Module UDP Round Trip Time

## Description

Measured the two way round trip time with packets of different sizes and
payload data.

Client sent at least two packets for each measurement class. The first
packet is always removed from measurement because DNS, path calculation,
route cache effects **SHOULD** not be measured.

The packet size **MAY** be larger then the minimum path MTU.

## Packet Behavior and Traffic Pattern

```
0.000: src -> [4 byte] -> dst      # not taken into measurement
0.010: src -> [4 byte] -> dst
0.020: src -> [4 byte] -> dst
[x8 packets]

0.x00: src -> [50 byte] -> dst   
0.x10: src -> [50 byte] -> dst
0.x20: src -> [50 byte] -> dst
[x8 packets]
```

This is redone for the following packet sizes:

`4, 50, 100, 250, 500, 1000, 1480`

## Packet Data

To determine packet loss and map the receiving packet to the sending packet a
sequence number must be added to the payload. This is a 4 byte network encoded
`uint32_t`. Starting with 0. The receiver **MUST** return the whole received packet
without any packet modification - including the remaining part

## Configuration Values

### Interval Time

Time between two transmitted packets. Default is 10 ms. The default interval
time is fine for MiB links, for links in kB field and lower bandwidth networks
this may lead to congestion and probably queue packet loss. This is not what is
intended with this module, so the administrator can select an higher interval
time.
