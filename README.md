Spwnnmark - use spwnn for benchmarking parallel machine performance

Various Results (go.1.18.6):
```
aws c7g.16xlarge                                            Elapsed time     6.121933222s; GOMAXPROCS 64
aws r6i.32xlarge                                            Elapsed time     7.193731471s; GOMAXPROCS 128
AMD Ryzen 7 1700X Eight-Core Processor 3.40 GHz             Elapsed time    44.689435500s; GOMAXPROCS 8
Intel 7700HQ                                                Elapsed time  1m24.816407300s; GOMAXPROCS 4
Intel(R) Core(TM) i5-4590 CPU @ 3.30GHz                     Elapsed time  1m31.449798500s; GOMAXPROCS 4
aws t4g.small (plenty of CPU credits)                       Elapsed time  3m06.854805956s; GOMAXPROCS 2
aws t3a.nano (plenty of CPU credits)                        Elapsed time  5m00.691448894s; GOMAXPROCS 2
Intel(R) N3350 2 cores @ 1.10GHz (Chromebook)               Elapsed time  8m31.797251343s; GOMAXPROCS 2
Intel(R) Atom(TM) Processor E3930 @ 1.30GHz                 Elapsed time 12m04.827775875s; GOMAXPROCS 2
raspberry pi 3  bullseye 64-bit golang 1.15                 Elapsed time 19m49.630674793s; GOMAXPROCS 4
raspberry pi 2b buster 32-bit & bullseye 32-bit             Elapsed time 20m16.867081494s; GOMAXPROCS 4
raspberry pi 3  buster 32-bit                               Elapsed time 23m22.680746919s; GOMAXPROCS 4
gcp e2-micro (you only get 20 second bursts then very slow) Elapsed time 42m59.583750118s; GOMAXPROCS 2
```

Amusingly the c7g.16xlarge is a Graviton 3 and costs a bit over $2.00 / hour.
The rci.32xlarge with 128 virtual cores (but really only 64 real ones and 64 useless hyperthreaded ones) is Intel and costs a bit over $8.00 / hour.
Go Graviton 3!

Why is the raspberry pi 2b faster than the 3 for 32-bit code?  I'm guessing it's because of 32-bit Go and 32-bit code optimization was the design goal of the 2b CPU.  (And there is no 64-bit Raspberry Pi OS for the 2b.)
