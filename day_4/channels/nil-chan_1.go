package main

import (
	"fmt"
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

/*
What happens when we receive from a closed channel?
We get the default value of the type of the channel.
In our case, the type is int so the value is 0.
*/
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
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
