Spwnnmark - use spwnn for benchmarking parallel machine performance

Results:
```
aws c7g.16xlarge                                            Elapsed time     6.121933222s; GOMAXPROCS 64
aws r6i.32xlarge                                            Elapsed time     7.193731471s; GOMAXPROCS 128
AMD Ryzen 7 1700X Eight-Core Processor 3.40 GHz             Elapsed time    44.689435500s; GOMAXPROCS 8
Intel 7700HQ                                                Elapsed time  1m24.816407300s; GOMAXPROCS 4
Intel(R) Core(TM) i5-4590 CPU @ 3.30GHz                     Elapsed time  1m31.449798500s; GOMAXPROCS 4
aws t4g.small (plenty of CPU credits)                       Elapsed time  3m06.854805956s; GOMAXPROCS 2
aws t3a.nano (plenty of CPU credits)                        Elapsed time  5m00.691448894s; GOMAXPROCS 2
raspberry pi 2b                                             Elapsed time 20m16.867081494s; GOMAXPROCS 4
raspberry pi 3                                              Elapsed time 23m22.680746919s; GOMAXPROCS 4
gcp e2-micro (you only get 20 second bursts then very slow) Elapsed time 42m59.583750118s; GOMAXPROCS 2
```
