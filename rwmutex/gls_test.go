package rwmutex

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func getSetDel(group *sync.WaitGroup) {
	for i := 4; i < 32; i++ {
		if i&3 == 0 { // multiple of 4
			Set(i)
			continue
		} else if i&(i+1) == 0 {
			Del()
		} else {
			val := Get().(int)
			if val >= i {
				panic("invalid")
			}
		}
	}
	group.Done()
}

func getSet(group *sync.WaitGroup) {
	for i := 4; i < 44; i++ {
		if i&3 == 0 { // multiple of 4
			Set(i)
			continue
		} else {
			val := Get().(int)
			if val >= i {
				panic("invalid")
			}
		}
	}
	group.Done()
}

func TestSet(t *testing.T) {
	wg := &sync.WaitGroup{}
	goRtns := 100 * 1000
	wg.Add(goRtns)
	for i := 0; i < goRtns; i++ {
		go getSetDel(wg)
	}
	wg.Wait()
}

func TestGet(t *testing.T) {
	val := Get()
	if val != nil {
		panic("invalid scenario")
	}
}

func BenchmarkSetGetDel(b *testing.B) {
	wg := &sync.WaitGroup{}
	goRtns := b.N
	wg.Add(goRtns)
	for i := 0; i < goRtns; i++ {
		go getSetDel(wg)
	}
	wg.Wait()
}

func BenchmarkSetGet(b *testing.B) {
	wg := &sync.WaitGroup{}
	goRtns := b.N
	wg.Add(goRtns)
	for i := 0; i < goRtns; i++ {
		go getSet(wg)
	}
	wg.Wait()
}

// TestSetGet1 validates memory not increasing over the time.
// It also prints total time it took for GetSet(10-Set,30-Get ops) for 10^5 go-routines.
func TestSetGet1(t *testing.T) {
	mStat := &runtime.MemStats{}
	goRtns := 100 * 1000
	for i := 0; i < 100; i++ {
		wg := &sync.WaitGroup{}
		st := time.Now()
		wg.Add(goRtns)
		for i := 0; i < goRtns; i++ {
			go getSet(wg)
		}
		wg.Wait()

		runtime.ReadMemStats(mStat)
		fmt.Println(time.Now().Sub(st), mStat.Alloc>>20)
		<-time.After(200 * time.Millisecond)
	}

}
