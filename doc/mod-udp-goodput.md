# Module UDP Goodput

## Description

Optimized for Goodput measurement. Packet loss and packet reordering is is not
focus of this module.

All UDP features of nuttcp should be supported, e.g.: `nuttcp -l8972 -T30 -u
-w4m -Ru -i1 <dest>`


## Features

- Payload pattern: zeroized, random ASCII, full rand
- Unicast and Multicast Support
- Parallel Workers (thread support)
- Goodput Target Bandwidth (user can select the nominal egress bandwidth)
- Burst Mode Support


## Configuration

```
{
  # per default one TCP transmitter is started, to spawn exactly
	# to much threads are cores are available use "cores".
	# Use "threads" if you want to fully utilize all virtual cores,
	# including hyperthreads.
	# If the system has several sockets, all sockets are utilized for
	# "cores" and "threads".
	"worker" : "1"

  # port for listening and sending. If worker is larger as 1 subsequent
	# ports are used. E.g. 7001, 7002, ...
	"port" : "7000"

	# payload pattern. Default is zeroized because we want to fullfill
	# the pipe and offload as much as possible. 
	"payload-pattern" : "zeroized"

	# limits the outgoing rate. Normally this is unlimited (value "0"): mapago
	# send as much data as possible without further configuration. The rate
	# can be given in any SI/IEC prefix form: 23mbit, 23mibit, ... just everything
	# as well it is unambiguous.
	# Note that rate depends on the "packet-length" parameter.
	"rate" : "0"

  # if rate is != 0 the rate-burst can be given. Normally the spacing between
	# packets is equal for a given calculated target rate. With burst given a burst
	# pattern can be given. These packets are transmitted without any pause.
	"rate-burst": "0"

	# The packet size to be send. The default is 512 byte, IPv4/IPv6 as well as UDP
	# is not considered. This is just the payload size. 512 byte is considered safe:
	# assume IPv4 the "minimum maximum reassembly" buffer size is 576 byte as specified
	# in RFC 1122. Minus IPv4 header (20) byte and UDP header (8) byte 512 is fine. Note
	# that due to IPv4 options the available payload can be smaller. But this is more
	# theoretical and 512 byte is fine.
	# To get line rate you probably want to increase this to jumbo mtu 9k/16k packet size.
	"packet-length" : "512"

  # set the DSCP value, unmodified will not modify the default
	"dscp" : "unmodified"

  # is OS default ttl
	"ttl" : "unmodified"

  # Can be human or json
	"output-format" : "human"

}
```

## Not Supported

- Packet loss and reordering detection
- Read payload from STDIN or from file

## Output Format

### Human


### JSON

The JSON format must be compatible between all peers. But not all Operating Systems
implement or provide the same functionality. Therefore the output format is splitted
into a mandatory and a optional part. All fields in the mandatory set must be available
for all compatible implementations.

```
{
  "core" : {

	},
	"aux" : {
	}
}

```
