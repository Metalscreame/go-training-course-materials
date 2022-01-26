package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Multiple functions can read from the same channel until
that channel is closed; this is called fan-out.
This provides a way to distribute work amongst a group of
workers to parallelize CPU use and I/O.

A function can read from multiple inputs
and proceed until all are closed by multiplexing
the input channels onto a single channel that's closed when
all the inputs are closed. This is called fan-in.

We can change our pipeline to run two instances of sq,
each reading from the same input channel.
introduce a new function, merge, to fan in the results:
 */

func main() {
	start := time.Now()
	in := gen(2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)
	//c3 := sq(in) // add 2 more to show time increase
	//c4 := sq(in)
	//c5 := sq(in)
	//c6 := sq(in)
	//c7 := sq(in)
	//c8 := sq(in)  c3, c4, c5, c6, c7, c8

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2,) {
		fmt.Println(n)
	}

	finish := time.Now()
	fmt.Println(finish.Sub(start))
}
// 1st will take data and will start putting it into channel, will return channel with results
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * 100)
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}