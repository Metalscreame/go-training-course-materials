package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	c := make(chan struct{})
	go func() {
		c <- struct{}{}
	}()
	go func() {
		_ = <-c
	}()

	<-c
}
