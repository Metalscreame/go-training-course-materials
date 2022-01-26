package main

import (
	"fmt"
	"time"
)
// this is pipeline pattern
func main() {
	start := time.Now()
	// Set up the pipeline.
	c := gen(2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15)
	out := sq(c)

	// Consume the output.
	for v := range out {
		fmt.Println(v)
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
