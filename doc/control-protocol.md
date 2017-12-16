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


## Message Sequences

### Info Request

| Field Name  | Required |
| ----------- | -------- |
| `id` | yes |

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
}
```


### 
