package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestCurRoutineID(t *testing.T) {
	wg := &sync.WaitGroup{}
	totalRtns := 1000
	wg.Add(totalRtns)
	for i := 0; i < totalRtns; i++ {
		go func(wg *sync.WaitGroup) {
			rid := CurRoutine()
			if rid == nil {
				panic(fmt.Sprintf("undesired go routine id %v", rid))
			}
			wg.Done()
		}(wg)
	}
	wg.Wait()
}
