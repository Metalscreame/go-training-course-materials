package main

import (
	"sync"
	"testing"
)

/*
The Go testing package contains a benchmarking facility
that can be used to examine the performance of your
Go code. This post explains how to use the testing
package to write a simple benchmark.
*/

// go test -bench=. -benchmem -run
var m = sync.Mutex{}

func BenchmarkMutexLockOutside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}

// go test -bench=. -benchmem -run
func BenchmarkMutexLockInside(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var m = sync.Mutex{}
		m.Lock()
		m.Unlock()
	}
}
func BenchmarkRWMutexLock(b *testing.B) {
	m := sync.RWMutex{}
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}

func BenchmarkRWMutexRLock(b *testing.B) {
	m := sync.RWMutex{}
	for i := 0; i < b.N; i++ {
		m.RLock()
		m.RUnlock()
	}
}

/*
As we can notice, read locking/unlocking a sync.RWMutex
is faster than locking/unlocking a sync.Mutex. On the other end,
calling Lock()/Unlock() on a sync.RWMutex is the slowest operation.

In conclusion, a sync.RWMutex should rather be used when we have
frequent reads and infrequent writes.
*/
