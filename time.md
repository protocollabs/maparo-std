# Maparo Time

All internally measured and transferred timevalue should use a realtime clock
(`CLOCK_REALTIME`). `CLOCK_MONOTONIC` principle be used if a clean synchronization
between client and server can de done. This is principle be true for remote mode.
But because the time synchronization is a) not that accurate as required and b) not
possible at all if operated in a non-remote mode.

The only solution is to ignore this within maparo. If a high resolution timing
analysis between client and server is required the only solution is to use GPS
mouses and disable gpsd for the time of measurement to do not risk NTP adjustments.

`CLOCK_MONOTONIC_RAW` would be fine if we can build upon maparo internal time
synchronization mechanisms - but we can't.

## Exchanged Time Format

If time is exchanged via JSON the format must be UTC and with a resulution
of microseconds:

```
2017-12-16T12:32:42.763987000
```

Implementations SHOULD check the number of digits of the fractions. If the number
is six then microseconds is used. If 9 digits it should be interpreted as nanoseconds.

With Python3:

```python
import datetime
dt = datetime.datetime.utcnow()
print(dt.strftime('%Y-%m-%dT%H:%M:%S.%f'))
```

For Golang:

```go
import "time"
import "fmt"
t = time.Now().UTC()
fmt.Println(t.Format("2006-01-02T15:04:05.000000000"))
```
