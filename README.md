Spwnnmark - use spwnn for benchmarking parallel machine performance

Results:
```
aws t3a.nano (plenty of CPU credits)    Elapsed time 5m0.691448894s; GOMAXPROCS 2
aws t4g.small (plenty of CPU credits)   Elapsed time 3m6.854805956s; GOMAXPROCS 2
aws c7g.16xlarge                        Elapsed time   6.121933222s; GOMAXPROCS 64
scw-lt                                  Elapsed time 1m24.81640730s; GOMAXPROCS 4

```