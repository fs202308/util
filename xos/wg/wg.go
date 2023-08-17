package wg

import (
	"sync"
)

//WaitGroup封装结构
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(f func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		f()
	}()
}

func (w *WaitGroupWrapper) WrapParam(f func(...interface{}), args ...interface{}) {
	w.Add(1)
	go func() {
		defer w.Done()
		f(args...)
	}()
}
