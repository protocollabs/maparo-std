### QUIC throughput

Name: `quic-throughput`

#### Description

`quic-throughput` completes Mapago to create QUIC streams.  

> Note: Mapago - a Go implementation of maparo - currently supports IETF draft-17 for QUIC. The used QUIC implementation is "quic-go", also written in Go (https://github.com/lucas-clemente/quic-go).

#### Features
- Parallel Workers (thread support)
- Call Size 

#### Name

This module is standardized with the name:

```
quic-throughput
```

#### Measurement Start Request

##### Info Reply

`quic-throughput` has no additional information for the client. The `quic-throughput`
dictionary MUST be empty.

E.g.

```
[]
  "id" : "hostname=uuid",
  "seq-rp" : <uint64_t>
  "modules" : {
     "quic-throughput" : {},
  }
[]
```

##### Measurement Start Request

```
{
	"streams" : "1"
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