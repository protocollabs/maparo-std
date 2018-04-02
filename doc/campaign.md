# Campaign

```
{
	"config" : {
		"addr" : "::1",
		"port" : "6666"
	},

	"mods" : [
		{
			"delay" : "rand(0.0, 1.0)",
			"module" : "udp-ping",
			"mode" : "client",
			"args" : [
				"addr=$addr",
				"port=$port"
			]
		},
		{
			"delay" : "0",
			"module" : "udp-ping",
			"mode" : "client",
			"args" : [
				"addr=$addr",
				"port=$port"
			]
		}
	]
}
```
