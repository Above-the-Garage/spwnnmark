Spwnnmark - use spwnn for benchmarking parallel machine performance

Results:
```
aws c7g.16xlarge                        Elapsed time     6.121933222s; GOMAXPROCS 64
scw-lt                                  Elapsed time  1m24.816407300s; GOMAXPROCS 4
aws t4g.small (plenty of CPU credits)   Elapsed time  3m06.854805956s; GOMAXPROCS 2
aws t3a.nano (plenty of CPU credits)    Elapsed time  5m00.691448894s; GOMAXPROCS 2
raspberry pi 3				Elapsed time 23m22.680746919s; GOMAXPROCS 4
```

