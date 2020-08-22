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
* Memory usage pattern over the time where goroutines(fix-numbered) continuously using GLS —> validates goroutine reuse design[Caveats-1], validates memory is not continuously increasing. 
* Long running(1 week) tests and verify memory usage //todo

##### Caveats:

* The solution is leveraging golang's goroutine reuse design hence any change around it may hamper the performance of this Library.
* Memory requirement: 
    * Best, Avg cases: ``2*pointer-size`` for every key inserted, Worst-case: ``4*pointer-size``

#### Supportability:

The library have been tested only on below architectures. However, it can work in other as well.
* go version: 1.9+
* arch: amd64
