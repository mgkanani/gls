Goroutine Local storage
----

Go as a language hides underneath complexities associated with a goroutine from intra-communication to blocking/unblocking it. It provides very minimal control over goroutines. Under these circumstances, goroutine local storage(GLS) pattern is tricky. In most of the scenarios, gls can be avoided. GLS pattern follows implicit communication mechanism and lead to code-complexities. However, in some cases GLS-benefits overweights its cons.

There have been certain push from many folks to provide GLS with standard golang packages. Later, core  golang developers decided not to provide GLS with standard golang packages.

This Library provides GLS with high performances. 
It leverages sync.Map underneath. After sometime, map stabilities and read-write are atomically performed without locking.

Some experiments have been carried out to verify its safety and more will be added:
* Obvious validating tests
* Memory usage pattern over the time where goroutines(fix-numbered) continuously using GLS —> validates goroutine reuse design[Caveats-1], validates memory is not continuously increasing. 
* Long running(1 week) tests and verify memory usage //todo

Caveats:
* The solution is leveraging golang's goroutine reuse design hence any change around it may hamper the performance of this Library.
* Memory requirement: 
    * Best, Avg cases: 2*pointer-size for every key inserted, Worst-case: 4*pointer-size
