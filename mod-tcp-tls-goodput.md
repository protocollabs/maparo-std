
### TCP TLS Goodput

Name: `tcp-tls-goodput`

#### Description

'tcp-tls-goodput' is identical to 'tcp-goodput', except it supports the TLS protocol, ensuring data privacy and data integrity.

> Note: Mapago - a Go implementation of maparo - currently supports TLS 1.2.

#### Features

- Payload pattern: zeroized, random ASCII, full rand (client side)
- Configurable DSCP value (if OS support, client side)
- Flexible traffic exchange configuration possibilies
- Setting No Delay
- Setting Maximum Segmet Size
- Goodput Limit configuration support. For TCP this is somewhat
  hacky. Especially for high data rates the userspace interaction can
  limit the overall system performance. So consider this feature as
  experimental
- Zerocopy mode uses snedfile (if OS support). Will not work in all
  other configuration options
- IPv6 flowlabel support
- Selectable congestion control algorithm
- Parallel Workers (thread support)

#### Name

This module is standardized with the name:

```
tcp-tls-goodput
```

#### Measurement Start Request

##### Info Reply

`tcp-tls-goodput` has no additional information for the client. The `tcp-tls-goodput`
dictionary MUST be empty.

E.g.

```
[]
  "id" : "hostname=uuid",
  "seq-rp" : <uint64_t>
  "modules" : {
     "tcp-tls-goodput" : { "cores" : <uint64_t> },
  }
[]
```

##### Measurement Start Request

```
{
	"streams" : <uint64_t>
}
```

#### Measurement Start Reply

```
{
  "streams" :
  [
		{ "listen-port" : "<port>" }
	],
}
```


#### Measurement Info Reply

```
{
  "streams" :
  [
	  {
		"timestamp-first" : "<maparo-time>"
		"timestamp-last"  : "<maparo-time>"
		"received-bytes"  : "<uint64_t>"
		}
	],
}
```

#### Not Supported

- Ignore <n> seconds from start of measurement. This must be done by analysis
	tooling