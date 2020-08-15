// +build go1.9

package gls

import (
	"github.com/mgkanani/goroutines"
	"runtime"
	"sync"
)

const factor = 10

var (
	shards = runtime.GOMAXPROCS(-1) * 2 // will be executed at load time.
	// most of the time, number of core are 2^x, but can be different due to virtualisation/containerisation
	// bitwise AND can be used when shards is 2^x.
	// division = shards - 1
	glsArr = make([]sync.Map, shards)
)

// Set accepts a value. key will be the current go-routine.
func Set(value interface{}) {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards
	glsArr[idx].Store(curRtn, value)
}

// Get returns a value present in map for calling go-routine
func Get() interface{} {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards
	val, ok := glsArr[idx].Load(curRtn)
	if !ok {
		return nil
	}
	return val
}

// Del deletes the value and key from map. Try to avoid this unless service is shutting down or pausing
// for pretty long. You can use Set(nil) instead.
func Del() {
	curRtn := goroutines.CurRoutine()
	idx := int(uintptr(curRtn)>>factor) % shards
	glsArr[idx].Delete(curRtn)
}
