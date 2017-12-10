# Shared Functionality

## Payload Pattern

Maparo pre-defines several payload pattern to be
used in each module.

The pattern is true for one "chunk". One chunk is one pre allocated buffer and
is typically one UDP packet or one large (max 4 GB) TCP chunk. Chunks are
reused and pattern are not recalculated -> identical. This is for performance
aspects because randomizing and touching chunks are CPU intensive and may lower
the network performance. I don't see any network measurement advantageous where
recalculating is required. Often optimizer and gzip for UDP work on a packet
level and for TCP the chunk size can be quite large so there is no real problem
with this limitation.

# Zero

Just 0 for the complete payload

Name: `zero`

### Random ASCII (letter)

Randomized string with `a-zA-Z0-9`. No Unicode

Name: `random-ascii`

### Random 

Random is a pure random byte.

Name: `random`

The generator tries to use the most cryptographic random bytes from the
underlying operating system (e.g. `/dev/random` seed combined with AES)


