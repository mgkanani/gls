package goroutines

import "unsafe"

type g struct{}

func goRoutine() *g

// CurRoutine returns current go routine pointer.
func CurRoutine() unsafe.Pointer {
	return unsafe.Pointer(goRoutine())
}
