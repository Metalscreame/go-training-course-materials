package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// from just for func
//https://www.youtube.com/watch?v=t9bEg2A4jsw
/*
the default value for channels is nil.
But not many of us know that this nil value is actually useful.

Given a nil channel c:
<-c receiving from c blocks forever
c <- v sending into c blocks forever
close(c) closing c panics
 */
func main() {
	a := sender(1, 3, 5, 7)
	b := sender(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}


func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					log.Println("a is done")
					adone = true
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					log.Println("b is done")
					bdone = true
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func sender(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}