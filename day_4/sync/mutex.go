package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Let’s take a look at the Go package in charge to provide synchronization primitives: sync.
First one is mutex
Let's imagine we have a map that should be update concurrently
*/

var m map[string]int

func set(k string, v int) {
	mux.Lock()
	defer mux.Unlock()

	m[k] = v
	fmt.Printf("Set %v\n", v)
}

func main() {
	m = make(map[string]int)
	go set("key1", 1)
	go set("key2", 2)
	go set("key3", 3)
	go set("key4", 4)

	// run few times and see panic
	// add mutex
	time.Sleep(time.Second)
}

var mux sync.Mutex
var rwmux sync.RWMutex

func get(k string) int {
	rwmux.RLock()
	defer rwmux.RUnlock()

	return m[k]
}

/*
sync.RWMutex
sync.RWMutex is a reader/writer mutex.
It provides the same methods that we have just seen with Lock()
and Unlock() (as both structures implement sync.Locker interface ).
Yet, it also allows concurrent reads using RLock() and RUnlock() methods:

A sync.RWMutex allows either at least one reader or exactly one writer whereas
a sync.Mutex allows exactly one reader or writer.

 */
