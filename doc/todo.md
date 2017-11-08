
maparo mod-udp-rtt --mode server --port 8888
maparo mod-udp-rtt --mode client --dst 1.1.1.1 --port 8888


maparo mod udp-rtt server config.port=8888
^^^^^^^^^^^^^^^^^^^^^^^^^
- This has a pre-defined config
- mods comes as pairs, for client and server, always


maparo mod-config udp-rtt server dst=10

To print udp-rtt server config as JSON file to STDOUT


maparo campaign-rrul server
maparo campaign-rrul client config.dst=::1

Campaign rrul client:

[ delay:0.0, mod:udp-rtt, config.dst=$DST ]
[ delay:1.0, mod:udp-goodput, config.dst=$DST ]



# Control Plane

maparo remote --key secret

# remote means that the other peer must operate in
# strict remote mode. E.g. the operation mode is
# dictated from the client. A remote process one client
# after another. Remote servers already processing a client
# all other incoming connects are dropped.
#
# The secret is just optional and is now has the same
# security level as the SNMP community string. Though,
# it is used as the encryption and prevents script kiddies
# to use a remote server somehow.
maparo --remote <ip:port> --remote-secret <secret> mod-udp-rtt --mode server --port 8888
