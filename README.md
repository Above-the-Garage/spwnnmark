Spwnnmark - use spwnn for benchmarking parallel machine performance

Various Results (go.1.18.6 unless otherwise noted):
```
aws c7g.16xlarge (Graviton3)                                Elapsed time       6.121933222s; GOMAXPROCS 64
aws m7i-24large (Intel)                                     Elapsed time       6.966997985s; GOMAXPROCS 96
aws r6i.32xlarge (Intel)                                    Elapsed time       7.193731471s; GOMAXPROCS 128
aws m7i.48xlarge (Intel)                                    Elapsed time       8.156488503s; GOMAXPROCS 192 (go 1.9.3)
Intel i9-12900H 2.50 GHz (6 HT p-cores; 8 e-cores)          Elapsed time      18.647013700s; GOMAXPROCS 20
AMD Ryzen 9 5900HX (HT)                                     Elapsed time      34.299816600s; GOMAXPROCS 16
AMD Ryzen 7 1700X Eight-Core Processor 3.40 GHz             Elapsed time      44.689435500s; GOMAXPROCS 8
t4g.2xlarge (Graviton2)                                     Elapsed time      49.708390402s; GOMAXPROCS 8
r7g.xlarge (Graviton3)                                      Elapsed time    1m06.859180419s; GOMAXPROCS 4
Intel 7700HQ                                                Elapsed time    1m24.816407300s; GOMAXPROCS 4
Intel(R) Core(TM) i5-4590 CPU @ 3.30GHz                     Elapsed time    1m31.449798500s; GOMAXPROCS 4
aws m7i-flex.xlarge                                         Elapsed time    1m32.732027200s; GOMAXPROCS 4
aws c5.xlarge                                               Elapsed time    2m07.145813700s; GOMAXPROCS 4
aws t4g.small (Graviton2, plenty of CPU credits)            Elapsed time    3m06.854805956s; GOMAXPROCS 2
aws c6i.xlarge (Intel)                                      Elapsed time    1m51.21501840s0; GOMAXPROCS 4
Intel N4120 @ 1.1GHz                                        Elapsed time    4m12.856900000s; GOMAXPROCS 4
aws t3a.nano (plenty of CPU credits)                        Elapsed time    5m00.691448894s; GOMAXPROCS 2
aws t3a.small (plenty of CPU credits)                       Elapsed time    5m09.107800250s; GOMAXPROCS 2
raspberry pi 4 B                                            Elapsed time    5m14.034612514s; GOMAXPROCS 4
Intel(R) N3350 2 cores @ 1.10GHz (Chromebook)               Elapsed time    8m31.797251343s; GOMAXPROCS 2
Intel(R) Atom(TM) Processor E3930 @ 1.30GHz                 Elapsed time   12m04.827775875s; GOMAXPROCS 2
raspberry pi 3  bullseye 64-bit golang 1.15                 Elapsed time   19m49.630674793s; GOMAXPROCS 4
raspberry pi 2b buster 32-bit & bullseye 32-bit             Elapsed time   20m16.867081494s; GOMAXPROCS 4
raspberry pi 3  buster 32-bit                               Elapsed time   23m22.680746919s; GOMAXPROCS 4
gcp e2-micro (you only get 20 second bursts then very slow) Elapsed time   42m59.583750118s; GOMAXPROCS 2
raspberry pi 3 w bullseye 32-bit go.1.8.4                   Elapsed time 1h03m46.311928563s; GOMAXPROCS 1
raspberry pi zero w bullseye 32-bit go 1.8.4)               Elapsed time 1h50m21.238283491s; GOMAXPROCS 1
```

Amusingly the c7g.16xlarge is a Graviton 3 and costs a bit over $2.00 / hour.

The t4g.2xlarge is pretty bad-ass Graviton - 8 real cores that an burst up for 9 hours.  You an get a lot done on that!  $200/month.
The rci.32xlarge with 128 virtual cores (but really only 64 real ones and 64 useless hyperthreaded ones) is Intel and costs a bit over $8.00 / hour.
Go Graviton 3!

The Intel N4120 is in a $250 Lenovo 2-in-1 laptop.  That level of performance is a bit shocking for such a cheap system.

Go Raspberry Pi 4 B

Why is the raspberry pi 2b faster than the 3 for 32-bit code?  I'm guessing it's because of 32-bit Go and 32-bit code optimization was the design goal of the 2b CPU.  (And there is no 64-bit Raspberry Pi OS for the 2b.)
