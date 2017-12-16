# Control Protocol

## Basics

Each server - started with argument `remote` listening on a TCP _and_ UDP port
for incoming control messages. The control protocol is fully optional, each
operation must be possible without a control protocol, though the program
arguments must be set manually and the result set must be merged manually by
using USB stick or some other transfer method.

## Unicast

## Multicast

If a remote server receives a UDP multicast request, the reply must be a UDP
unicast.  The unicast reply must address the sending IPv{4,6} address.


### Protocol Requirements

The very first four bytes of a packet must encode the control protocol type.
The type has an associated encoding format (i.e. JSON). But this can be
different.

### Protocol Types

In `uint32_t`, network byte order:

- `0`: null message request
- `1`: null message reply
- `2`: info request
- `3`: info reply


## Messages

### Null Messages

Null Messages are noop messages with no direct usage beside "warming the
cache". Why info-request, info-reply messages measure the round trip times.
To be accurate the cell phone should not be in deep idle mode, systemd should
have started the daemon, the routing cache should be filed, IPv4 APR and IPv6
neighbor discovery mechanisms should be warmed. etc.

The purpose of NULL messages is to warm up the pipe.

Subsequent messages (e.g. info-request) SHOULD be transmitted if the sending
and receive phase is finished. Only after the null-reply is received the
complete chain is processed (warmed).

The null message data do not matter. The sender is free to send binary or ascii
data. The reply is never touched, modified or checked in any way.

#### Null Request

The message size MUST NOT be larger as 512 bytes

#### Null Reply

The reply host SHOULD implement a ratelimiting component.
The reply host MUST transfer the data back to the sender as fast as possible.

The server is free to ignore payloads larger as 512 bytes.

### Info Message


#### Info Request

| Field Name  | Required |
| ----------- | -------- |
| `id` | yes |
| `seq` | yes |
| `ts` | no (optional) |

Generated from client, sent to TCP unicast address or UDP multicast
address if it is a multicast module or unicast if UDP unicast analysis.

```
{
  # to identify the sender uniquely a identifier must be transmited.
  # The id consits of two parts:
  # - a human usabkle part, like hostname or ip address if no hostname
  #   is available.
  # - a uuid to guarantee a unique name
  # Both parts are divided by "=", if the character "=" is in the human
  # part it must be replaced by something else.
  # The id is stable for process lifetime. It is ok when the uuid is 
  # re-generated at program start
  "id" : "hostname=uuid",

  # a sender may send several request in a row. To address the right one
  # the reply host will refelct the sequence number.
  # The sequence number should start with 0 for the first generated packet
  # but can start randomly too. The sequence number MUST be incremented at
  # at each transmission. In the case of an overflow the next sequence numner
  # MUST be 0. Strict unsigned integer arithmetic
  "seq" : <uint64_t>

  # The timestamp is replied untouched by the server. The timestamp can
  # be used by the client to calculate the round trip time.
  # The timestamp in maparo format with nanoseconds, optional
  # In UTC
  # format example: 2017-05-14T23:55:00.123456789Z
  "ts" : "<TS>"
}
```



#### Info Reply

| Field Name  | Required |
| ----------- | -------- |
| `id` | yes |

Generated from server, sent to TCP unicast address or UDP unicast
address. The address is the sender ip address.

Info messages should be replied as fast a possible. This is required to calculate
a clean round trip time. The info client SHOULD calculate as much as possible values
before the reception of info-request messages. I.e. the id can be calculated at
program start for example.

```
{
  # The Id identify the reply node uniquely. The id is generated in indentical
  # way as the info-request id.
  "id" : "hostname=uuid",

  # the RePlied sequence number from the sender
  "seq-rp" : <uint64_t>

  # the RePlied sequence number from the sender - if available. If not
  # nothing MUST be replied.
  "ts-rp" : "<TS>"

  # the timestamp in standard maparo format (also UTC). The
  # timestamp can be used to check (simplified) the time delta between
  # client and server. The client can warn the user if the times are not
  # synchronized or use the method to synchronize the time with the client
  # e.g. saving the calculated offset (neglect rtt and processing delays).
  # The first check if the replied timestamp is between the client info-request
  # sending timestamp and the client info-reply receive timestamp.
  "ts" : "<TS>"

  # Valid values:
  # - amd64
  # - 386
  # - arm
  # - arm64
  # - ppc64le
  # - s390x
  # - unknown
  "arch" : <ARCH>

  # valid values:
  # - linux
  # - windows
  # - freebsd
  # - osx
  # - android
  # - ios
  # - unknown
  "os" : <OS>
}
```


![image](images/control-time-measurement.svg)


### 
