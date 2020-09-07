Goroutine Local storage
----
[![Build Status](https://travis-ci.com/mgkanani/gls.svg?branch=master)](https://travis-ci.com/mgkanani/gls)
[![Coverage Status](https://coveralls.io/repos/github/mgkanani/gls/badge.svg?branch=master)](https://coveralls.io/github/mgkanani/gls?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mgkanani/gls)](https://goreportcard.com/report/github.com/mgkanani/gls)
[![GoDoc](https://img.shields.io/badge/Godoc-reference-blue.svg)](https://godoc.org/github.com/mgkanani/gls)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/mgkanani/gls)](https://pkg.go.dev/github.com/mgkanani/gls)

This Library is implementation of Thread Local Storage pattern in golang.
Go as a language hides underneath complexities associated with a goroutine from intra-communication(through channels)
 to blocking/unblocking it.
Unlike Java, golang provides very minimal control over goroutines. Under these circumstances, goroutine local storage(GLS) pattern is tricky.
In most of the scenarios, gls can be avoided. GLS pattern follows implicit communication mechanism and can lead to code-complexities.
However, in some cases GLS-benefits overweight their cons.

There have been certain push from many folks to provide GLS with standard golang packages.
Later, core golang developers decided not to provide GLS with standard golang packages.

#### Under the hood

This Library provides GLS with high performances. Over the period, access to gls-backed value is lock-less.
This library is built using: 
* ``sync.Map`` to store value for go-routines
* Scheduler's go-routines re-usability 
* Sharding to avoid performance bottleneck in `Map.Store` initially

After sometime, no new goroutines are created but scheduler uses existing one.
map stabilises and read-write are atomically performed without locking.

Below tests/experiments have been carried out to verify its safety and more can be added:
* Obvious validating and coverage tests
* Memory usage pattern over the time where goroutines(fix-numbered) continuously using GLS —> validates goroutine reuse design(i.e. Caveats-1), validates memory is not continuously increasing. 
* Long running(1 week) tests and verify memory usage //todo

##### Caveats:

* The solution is leveraging golang's goroutine reuse design hence any change around it may hamper the performance of this Library.
* Memory requirement: 
    * Best, Avg cases: ``2*pointer-size`` for every key inserted, Worst-case: ``4*pointer-size``


##### Comparison with Read-Write Mutex:
This comparison is more between sync.Map vs sync.RWMutex backed map for gls pattern with below scenario:
 * At the end of goroutine del() is being called ```TestGoRtnReUsageStatsWithDel```
 * At the end of goroutine del() is not called ```TestGoRtnReUsageStatsWithoutDel```

At the start of every iteration, 100K go-routines spawned. Each go routine performs 40 Get:Set in ratio 3:1 operations. There are 50 such iterations.
 
System configuration:
* 2 cores (core 2 duo)
* ~700MB free out of 2GB before below run

```
$ go test -v ./...

...
=== RUN   TestGoRtnReUsageStatsWithoutDel
Min:740.124215ms, Max:2.37434027s, Avg:1.103566728s
time taken:  55.178477167s
--- PASS: TestGoRtnReUsageStatsWithoutDel (57.87s)
=== RUN   TestGoRtnReUsageStatsWithDel
Min:737.743745ms, Max:2.134757081s, Avg:1.238451248s
time taken:  1m1.923180247s
--- PASS: TestGoRtnReUsageStatsWithDel (62.93s)
PASS
ok  	_/home/mahendra/gls/gls	124.513s

...
=== RUN   TestGoRtnReUsageStatsWithoutDel
Min:1.410195319s, Max:5.135148053s, Avg:3.359963156s
time taken:  2m47.998293153s
--- PASS: TestGoRtnReUsageStatsWithoutDel (172.19s)
=== RUN   TestGoRtnReUsageStatsWithDel
Min:1.73240396s, Max:3.50120815s, Avg:2.423918734s
time taken:  2m1.196079057s
--- PASS: TestGoRtnReUsageStatsWithDel (123.50s)
PASS
ok  	_/home/mahendra/gls/gls/rwmutex	302.087s
```

From above data, it is clear that gls is performing best and takes less than half-time compared to ReadWriteMutex based approach.

#### Supportability:

The library have been tested only on below architectures. However, it can work in other as well.
* go version: 1.9+
* arch: amd64
