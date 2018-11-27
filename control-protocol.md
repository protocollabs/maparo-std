# Maparo Control Protocol Definition

## Essential Characteristics

Each server - started with argument `remote` listening on a TCP _and_ UDP port
for incoming control messages. The control protocol is fully optional, each
operation must be possible without a control protocol, though the program
arguments must be set manually and the result set must be merged manually by
using USB stick or some other transfer method.

The default control port for TCP and UDP is 64321. The control port can be
adjusted to any other port. The control *should* listen to unicast and
multicast and bind to the wildcard address. Supporting IPv4 and IPv6. The
network protocol *should* be selectable to disable IPv4/IPv6 if required.


To protect public servers the control must support a secret mechanism. It
is intended not a cryptographic mechanism because multicast and cryptographic
is somehow hard to do. Also for unicast encryption with a TLS like mechanism
certificates must be setup which is not always possible and increase the
complexity. Maparo is a control application and used in walled environments.
The secret mechanism is comparable to SNMPv1 with the community string - not
more.

The server must response to each message - unique identified by the sequence
number - exactly once. The server MUST not response multiple times to one
sequence number. The client MUST NOT reuse the same sequence number again, the
sequence number must always be incremented by the client at each transmission.

> To increase robustness for lossy links the client may send several requests
> with increasing sequence number. The server *should* drop packets with already
> processed sequence numbers.

The sequence number serves as a duplicate and reobustness method within the control
packet sequence. To differentiate two ongoing, parallel measurements the
sequence number is not suffiently. To identify a measurment uniquely a "measurment-id"
is required.

> Use case: one client, one server setup. The server start with one UDP goodput
> measurement and one TCP goodput measurement in parallel. To get info from one
> particular module the client must identify the particular measurement. The
> identification of the module like "udp-goodput" is not sufficiently because
> two udp-goodput modules may operates on the same time. Therefor a 'measurment-id'
> was introduced.

The Control Protocol is optional. All implementations are engaged to implement
a mechanism on server and client side to use the same functionality without
the protocol requirements.

The control protocol is designed to work on top of UDP and TCP. Additional
for UDP the protocol is also designed from the ground up to operate via Multicast.

UDP for discovery is fine, but for control communication a client/server should
prefer TCP for reliable communication - if any possible. Normally a UDP based
communication for standard request/reply flow is fine, at least when collected
data must be transfered back from server to client at measurement stop and if
the collected data is larger as MTU sized packets a reliable control channel is
required. This can be done with UDP and implement all the fancy stuff, at the
and what is implemented looks like TCP - why not take TCP for all control activity?

If TCP is selected as control protocol the control connection SHOULD stay open
all the time. This is required to allow the server to send asynchonous messages
during the measurement. This is required for modules inmplementing a reverse
transmission (data transmission from server to client) and inform the client when
finished. The client cannot know this and may poll the server otherwise. Thus,
it is helpful when the TCP control connection stays open during the complete
measurment.

## Golden Rule of Operation

The Control Protocol MUST never influence the measurement in any way. For
example: during a TCP measurement the control protocol must absolutely
do nothing - no transmission at all. This is especially important if
test are done in environments with only several kb bandwidth.

The only exceptions are explicit switches where the user is explicitly
informed that control traffic is not send over the wire. Use cases where
permanent protocol exchange is required are progress bars where status
(transferred bytes) are updated live at client side, without waiting
until the transmission is ready.

## Unicast

For Unicast measurments the control protocol SHOULD use TCP - even if the
measurement protocol is UDP. If exact round trip time measurements are required,
the TCP timeouts has negative impact or if UDP has other advantages compared to
TCP, UDP can be used as the control protocol. Though, packet loss, reordering
must be handled by the control plane.

## Multicast

If a remote server receives a UDP multicast request, the reply must be a UDP
unicast.  The unicast reply must address the sending IPv{4,6} address.

It is possible that after a certain discovery phase (most likely INFO-REQUEST,
INFO-REPLY) and the "most wanted" server is selected the addressing change
from multicast to unicast adressing.

> There is no need to send Time Diff request/reply probes to an multicast address
> and filter the results later if several servers are within the multicast
> domain.

## Control Address and Data Address

Beside iperf and other performance measurement programs maparo splits
control and data channel for maximum flexibility. Most often the control
and data channels are routed over the same protocol and path. Sometimes
the setup requires something special. Imagine a network with loss of
50%, a TCP control channel will not work in such environments. To analyse
such networks it is required to provide two networks: a test network and
a control network. To support such environments maparo must differentiate
control and data plane.

> Example: **Maparo Pulser**
>
> Two options are available
> - addr
> - ctrl-addr
>
> If no addr is given (None), the addr can be derived from info-reply
> message originator addresses. This is an implementation detail.

If it is a multicast measurement the addr must be given. The ctrl-addr
must be a multicast address to. It SHOULD be the identical multicast address.

To discover maparo servers the ctrl address can be a multicast addresses.
To discover both IPv4 and IPv6 only hosts the control address can be a list.
E.g. `--ctrl-addr FF02::1,224.0.0.1`

If a data address is given the address has precedence and MUST overwrite
the control address if the control address is a multicast address. If the
control address is a unicast address both addresses MUST be untouched.

If no control address given (None) the application SHOULD take the data
address.

> This is the standard behavior and is what the user expect! In 90% of
> all use cases the measurment and control network is identical. The user
> should not be enforced to specify the same address for unicast and
> multicast twice.

If no data address is given it should auto discover the data address
by using discovery service by control address.

If no data and no control address is given the application should give
up. Alternatively the application can use `::1` or `127.0.01`. Although
it is unlikely that a user what this.


## Control Message Ordering and Sessions

Maparo Control Protocol is stateles - control session do not exist. There
are also no message order requirements. Clients are free to send whatever
messages they like. For example: a client can start with a Time Diff message
followed by a INFO info or vice versa.

The only "light" exception are module-start and module-stop messages. If
module-stop messages are transmitted before module-start the server cannot
answer corretly and will return a failure. But this is *not* handled within
the control protocol due to session idenfiers, it is handled within the
server exclusivly based on internal states.

## Reply Requirements & Behavior

A server **MUST** not answer to a client request. The behavior is not standardized
and open to implementers. Servers can use message type 255 to signal an generic
error condition.

Several possibilities why a server do now answer:

- do not implement the ctrl protocol itself (remember, ctrl protocol is optional)
- the server is bussy under a other measurement and has no cpu time left to answer
  another ctrl request. 

A server **SHOULD** answer with a ctrl message if something is broken or an ongoing
measurement is active.

> There is explicetly no hard requirement that a ongoing measurement blocks other
> measurement attempts. Implementations are free to implement from allowing parallel measurements
> to only one measurment with a negative warning/error message signaled back to
> the requester.

A implementation may lock a measurement between `module-start` and `module-stop`
sequence. Between these the real measurement take place. Control measurements
may not be locked in any way to reduce contention.

Clients should implement a backoff for `module-start` requests to prevent storms.

## Discovery Process and Dual Hosts Handling


![image](images/control-mcast-discovery-dual-stack.svg)

## Binary Encoded Header

The standard control protocol header is componsed of the following elements
for **all** protocol headers:

![image](images/control-header.svg)

- 2 byte, `uint16_t`, unsigned, network byte order encoded **type**
- 2 byte, n**reserved**
- 4 byte, `uint32_t`, unsigned, network byte order encoded packet **length**

> The reserved field can be used later to signal LZMA compressed payload
> or for enrypted control 



### Protocol Requirements

The very first four bytes of a packet must encode the control protocol type.
The type has an associated encoding format (i.e. JSON). But this can be
different.

### Protocol Types

In `uint32_t`, network byte order, starting with 1, 0 is intentionally left
blank:

- `1`: info request
- `2`: info reply
- `3`: measurement start request
- `4`: measurement start reply
- `5`: measurement stop request
- `6`: measurement stop reply
- `7`: measurement info request
- `8`: measurement info reply
- `9`: time-diff request
- `10`: time-diff reply
- `255`: warning and error message


## Messages


![image](images/control-example.svg)

### Time Diff Messages

The first 4 bytes of the payload contains a network byte order encoded length of
the payload len, not including this "header".

The first Time Diff message CAN be ignored to bypass measurement jitter because of unwarmed
caches, arp/nd setup, xinitd init sequences and other effects.

The client can use Time Diff message several times to increase the precicion of
measurements.

#### Time Diff Request

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
  # the reply host will send back the sequence number.
  #
  # A receiver MUST answer to one equest exactly once.
  #
  # Sequence numbers are message specific. For example: info request message
  # numbers start with 0, later module-start-request first packet also has
  # sequence number 0.
  #
  # The sequence number should start with 0 for the first generated packet
  # but can start randomly too. The sequence number SHOULD be incremented at
  # at each transmission. It is possible that the sequence number is not a
  # sequence, but is MUST guaranteed that a sequence number is not transmitted
  # twice. The trivial implementation is to transmit ordered.
  # In the case of an overflow the next sequence numner
  # MUST be 0. Strict unsigned integer arithmetic.
  # The value must be converted to string, this is required to align all
  # json encoding to string values everywhere. "seq" : "1" not "seq" : 1
  "seq" : <uint64_t>

  # The timestamp is replied untouched by the server. The timestamp can
  # be used by the client to calculate the round trip time.
  # The timestamp in maparo format with nanoseconds, optional
  # In UTC
  # format example: 2017-05-14T23:55:00.123456789Z
  "ts" : "<TS>"

  # to fill the data packet the client can use the padding field to inject
  # arbitrary data into the packet.
  #
  # The field is optional
  #
  # If not otherwise specific the padding data SHOULD be replied
  "padding" : <string>
 
  # if server requires a string the string is required.
  "secret" : <string>
}
```


#### Time Diff Reply

The reply host CAN implement a ratelimiting component.

The reply host MUST transfer the data back to the sender as fast as possible.

The server is free to ignore payloads larger as MTU sized packets bytes.

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

	# A sender may send several request in a row.  # A receiver MUST answer to
  # one equest exactly once.
  #
  # Sequence numbers are message specific. For example: info request message
  # numbers start with 0, later module-start-request first packet also has
  # sequence number 0.
  #
  # The sequence number should start with 0 for the first generated packet
  # but can start randomly too. The sequence number MUST be incremented at
  # at each transmission. In the case of an overflow the next sequence numner
  # MUST be 0. Strict unsigned integer arithmetic.
  # The value must be converted to string, this is required to align all
  # json encoding to string values everywhere. "seq" : "1" not "seq" : 1
  "seq" : <uint64_t>

  # The timestamp is replied untouched by the server. The timestamp can
  # be used by the client to calculate the round trip time.
  # The timestamp in maparo format with nanoseconds, optional
  # In UTC
  # format example: 2017-05-14T23:55:00.123456789Z
  "ts" : "<TS>"

  # to implement a trivial access mechanism a secret can be given.
  # if the server do not accept the string the request is dropped
  # and a warning should be printed at server side that the secret
  # do not match the expections.
  # If the server has no configured secret but the client sent a
  # secret, then the server SHOULD accept the request.
  "secret" : <string>
}
```



#### Info Reply

| Field Name            | Required |
| --------------------- | -------- |
| `id`                  | yes |
| `seq-rq`              | yes |
| `modules`             | yes |
| `control-protocol`    | no  |
| `arch`                | no  |
| `os`                  | no  |
| `info`                | no  |

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

  # the RePlied sequence number from the sender, the number is encoded as as
  # string. (E.g. "seq-rp" : "392192")
  "seq-rp" : "<uint64_t>"

  # list of supported modules, the entries must point to empty dictionaries
  # for now. Later the empty dictionaries can be filled if additional
  # information is required. The module specific info within the dictionary is
  # specified in the particular module specification. E.g. the "tcp-goodput"
  # module has a section which specifies allowed values for the tcp-goodput
  # module.
  "modules" : {
     "udp-goodput" : { },
     "tcp-goodput" : { },
  }

  # The info-reply SHOULD return a "control-protocol" block where all supported
  # control protocols are listed. The only supported attribute is port for now.
  # The internet protocol (IPv4, IPv6) is not covered YET. Later it can probably
  # a IP list be specified. The control protocol should be extended later if a
  # certain maturity level is reached.
  # The following examples illustrates a server with support for UDP, TCP and UDP
  # UDP Multicast. Each listening on port 64321 for control messages. If IPv4 and
  # IPv6 is supported and on which unicast address is not specified.
  "control-protocol" : {
        # supported control transport protocols and info
	      "transport" : {
						"udp" : {
								"port" : "64321",
						}
						"udp-mcast" : {
								"port" : "64321",
						}
						"tcp" : {
								"port" : "64321",
						}
        }
        # the server "reacts" to the following control message types:
        "message-types" : [
            "info", "measurement-start", "measurement-stop", "measurement-info", "time-diff"
        ]
  }

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

  # The server can reply a string where server specific information
  # can be held. Like banner information or implementation name.
  # The client SHOULD print this information to the user.
  # The info string MUST not larger as 32 bytes
  "info" : <string>
}
```

##### Time Offset Calculation and Visualization

![image](images/control-time-measurement.svg)


### Measurement Start

Used to start module on server.

The Measurement-start message is self-contained. All server actions depends on this
message and are stateless. There is no need for the server to store information
from previous time-diff-request, info-request or any other messages. This behavior
is intended.

#### Measurement Start Request

| Field Name  | Required |
| ----------- | -------- |
| `id` | yes |
| `seq` | yes |
| `measurement-id` | yes |
| `measurement` | yes |
| `secret` | no |
| `measurement-delay` | no |
| `measurement-time-max` | no |

```
{
  # The Id identify the reply node uniquely. The id is generated in indentical
  # way as the info-request id.
  "id" : "hostname=uuid",

  # a sequence to identify the answer. For UDP within a high loss environment
  # the client may send several requests. The server SHOULD never reply twice
  # or even more.
  "seq" : <uint64_t>

  # Randomly picked measurement id, stable for one measurment. All
  # subsequent requsts/replies to the particular message must use this measurment-id.
  # Think about two parallel ongoing udp-goodput measurements. The client will
  # alternating query info-reply messages for both ongoing measurements with
  # both measurement-id's.
  #
  # The client dicatates the measurement id for one measurements. The server will
  # reply this id in the reply message too.
  "measurement-id" : "<uint64_t>"

  # to implement a trivial access mechanism a secret can be given.
  # if the server do not accept the string the request is dropped
  # and a warning should be printed at server side that the secret
  # do not match the expections.
  # If the server has no configured secret but the client sent a
  # secret, then the server SHOULD accept the request.
  "secret" : <string>

  # if the measurment should start delayed a value in seconds
  # can be given. This is only usefull where the server starts
  # the measurment action (e.g. sending data from server to client)
  # and where probably multiple servers should starts somehow
  # synchronously.
  # If the server starts several measurment-start requests to
  # increase robustness the measrument-delay time must be adjusted
  # by the client.
  # If later a absolute time is required a "measurement-delay-time"
  # parameter can be added.
  "measurement-delay" : <uint32_t>

  # seconds after which the measurement is guaranteed not active
  # and finished. The server can close all resourches allocated
  # at measurement start time like open sockets, etc.
  # Normally a measurment-stop command frees all resourches at
  # server side. But UDP multicast setups in packet loss environments
  # the stop may get lost. The client is only able to estimate how
  # long a measurment will be.
  # If nothing is specified the default should be 5 minutes.
  # The server may - also depending on the actual measurment -
  # adjust the time maximum.
  # A server is free to reject measurment-time-max values out
  # out scope. E.g. if a user want to block a server for 20 minutes
  # or so. (see "secret" for a better option")
  # "measurement-time-max" is started after "measurement-delay" is 0.
  # Or in other words: after the measurment is actual started.
  # If TCP is used for the control protocol the allocated ressourches
  # can be deallocated if the TCP connection is closed/interrupted.
  # At the end: a maparo server operated in the internet should behave
  # save and self-healing under all circumstances: lost UDP control messages,
  # closed TCP control connections.
  "measurement-time-max" : <uint32_t>

  # the module specific configuration
  "data" = {
       # see protocol specific section, e.g.
       # mod-tcp-goodput.md and "Measurement Start Request"
  }
}
```

#### Measurement Start Reply

A server MUST answer to a measurement-start request with a measurement start reply
*after* all systems are started and ready to serve the client. A server MUST NOT
start the subsystems afterwards.

Background:

- if the server proactively answers with a reply and the ports are not started and
  the client send immediately a message the message will be lost. So everything
  must be setup before the reply message is transmitted to the client
- if at server side something fail, the server can send the error back to the
  client and inform the client. This is not possible if the answer is send
  immediately.

If the server do not receive any packets within a predefined duration the server
SHOULD assume that the client crashes and SHOULD restart to a sane state so that
other clients are able to connect and use the service.

> This can be implemented by spanning a timer and if within n minutes no packages
> arrived the server should cancel the state and switch to the initial state.

After a measurment was started and additional measurment-start are received the
server MUST react in the following manner:

- if the measurment-start-reqest was sent from the same <id> and the identical
  measurment was requested the server must answer with measurment-start-reply
  with status "ok"
- if the measurment is from the another client instance or the measurement
  is another the server should return with status code busy.

```
{
  # The Id identify the reply node uniquely. The id is generated in indentical
  # way as the info-request id.
  "id" : "hostname=uuid",

  # The replied measurement id from the server.
  "measurement-id" : <uint64_t>

  # the status of the previous request, can be (lowercase)
  # - "ok"
  # - "busy" if another measurement is ongoing and no capacity is available to
  #          start a new measurement. The client CAN automatically (backoff) come
  #          back to request a new module-start measurment.
  # - "warn" if start was sucessfull BUT there not all parameter can be fullfilled
  #          then warn can be used to signal such a condition
  # - "failed" if the measurement cannot be started. Another usage: if the measurement-id
	#            is already in use by this client. E.g. the client do not calculate a new
	#            measurement-id, the server will response with a "failed" state too.
  "status" : "<status>"

  # a human readable error message WHY it failed. Can be
  # missing. If status is != ok the message SHOULD be set. Normally
  # this message SHOULD be printed to the user. The client react to the
  # status, but the human can interpret the message and take appropiate
  # steps. E.g. inform the maparo server owner.
  "message" : "<string>"

  "seq-rp" : "<uint64_t>"

  # the module specific return value
  "data" = {
        # see protocol specific section, e.g.
				# mod-tcp-goodput.md and "Measurement Start Reply"
  }
}
```

### Measurement Info

Measurment info messages are used to gather measurement data during an
ongoing measurement without stoping the active measurement.

#### Measurement Info Request

```
{
  "id" : "hostname=uuid",
  "seq" : <uint64_t>

  # The measurment id where the info is gattered.
  "measurement-id" : <uint64_t>

  "secret" : <string>
}
```

#### Measurement Info Reply

The server SHOULD only return measurement info if the id is identical
to the measurement-start id. I.e. no other client should be able
to get live measurement data.

```
{
  "id" : "hostname=uuid",

  # The measurment id where the info is gattered.
  "measurement-id" : <uint64_t>

  "seq-rp" : <uint64_t>

  # the module specific configuration, see module specification (e.g.
  # tcp-goodput for one example)
  "data" = {
        # see protocol specific section, e.g.
				# mod-tcp-goodput.md and "Measurement Info Reply"
  }
}
```

### Measurement Stop

#### Measurement Stop Request

The module stop-request must be from the identical start-request sender. The
source IP doesn't matter. The id is important. The server MUST ignore packages
from other hosts sending a stop-request message. The server SHOULD print a warning
message on the console.

The server SHOULD implement a guard time after which the server should accept a
new module-start sequence.


```
{
  "id" : "hostname=uuid",
  # The measurment id where the info is gattered.
  "measurement-id" : <uint64_t>
  "seq" : <uint64_t>
  "secret" : <string>
}
```

#### Measurement Stop Reply

Most important rule: the server don't know when the measurement is over! Only
the client knows this information. When a measurement is over is dictated by
the client.


```
{
  "id" : "hostname=uuid",

  # the status of the previous request, can be (lowercase)
  # - "ok"
  # - "busy" if another measurement is ongoing
  "status" : <status>

  # The measurment id where the info is gattered.
  "measurement-id" : <uint64_t>

  "seq-rp" : <uint64_t>

  # the module specific configuration
  "measurement" = {
        "name" : <module-name>
        # the output data
        "data" : {
       }
  }
}
```
