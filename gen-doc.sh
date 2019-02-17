#!/bin/sh

rm -rf maparo-spec.md
cat README.md >> maparo-spec.md
cat control-protocol.md >> maparo-spec.md
cat command-line-interface.md >> maparo-spec.md
cat mod-tcp-goodput.md >> maparo-spec.md
cat mod-udp-goodput.md >> maparo-spec.md
cat mod-udp-ping.md >> maparo-spec.md
cat mod-udp-rtt.md >> maparo-spec.md
cat time.md >> maparo-spec.md

# remove svg for now, later start inkscape and
# convert to png
sed -i '/.svg/d' maparo-spec.md

VERSION=$(git describe --always --dirty)
ABSTRACT='A network performance measurement protocol specification. Beside iperf, netperf and other tools it just defines the protocol specification - not one particular implementation. Similar to HTTP/2 (RFC 7540) or any other networking protocol specification.'

pandoc maparo-spec.md --toc --pdf-engine=xelatex -V abstract="$ABSTRACT"  -V version="$VERSION"  -V title="Maparo -- A Network Performance Measurement Protocol Specification"  --standalone --template assets/technical.tex -o maparo-spec.pdf

rm -f maparo-spec.md
