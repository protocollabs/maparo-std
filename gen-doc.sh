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

pandoc maparo-spec.md --toc --pdf-engine=xelatex -o maparo-spec.pdf

rm -f maparo-spec.md
