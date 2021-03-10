package rwmutex

import (
	"github.com/mgkanani/gls/goroutines"
	"runtime"
	"sync"
	"unsafe"
)

const factor = 10

type Map struct {
	m map[unsafe.Pointer]interface{}
	sync.RWMutex
}

var (
	shards = runtime.GOMAXPROCS(-1) * 2 // will be executed at load time.
	// most of the time, number of core are 2^x, but can be different due to virtualisation/containerisation
	// bitwise AND can be used when shards is 2^x.
	// division = shards - 1
	glsArr = make([]*Map, shards)
)

func init() {
	for idx := range glsArr {
		glsArr[idx] = &Map{
			m: make(map[unsafe.Pointer]interface{}),
		}
	}
}

// Set accepts a value. key will be the current go-routine.
func Set(value interface{}) {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards
	glsArr[idx].Lock()
	glsArr[idx].m[curRtn] = value
	glsArr[idx].Unlock()
}

// Get returns a value present in map for calling go-routine
func Get() interface{} {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards
	glsArr[idx].RLock()
	defer glsArr[idx].RUnlock()
	val, ok := glsArr[idx].m[curRtn]
	if !ok {
		return nil
	}
	return val
}

// Del deletes the value and key from map.
func Del() {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards

	glsArr[idx].Lock()
	delete(glsArr[idx].m, curRtn)
	glsArr[idx].Unlock()
}
