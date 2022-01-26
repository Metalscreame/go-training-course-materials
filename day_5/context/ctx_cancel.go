package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cancel the context after two seconds
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(time.Second * 2)
		fmt.Println("Cancelling context")
		cancelFunc()
	}(cancel)


	go doWork(ctx, 3)
	go iterateOverSomething(ctx, 8)

	fmt.Scanln()
}

func doWork(ctx context.Context, timeSleep time.Duration) {
	fmt.Println("Starting doSomething")
	select {
	// Here I'm simulating the time doSomething will take to process.
	case <-time.After(time.Second * timeSleep):
		fmt.Println("Finished doWork")
	// Check the context Done to get cancellation signal
	case <-ctx.Done():
		fmt.Println("Cancelling do work")
	}
}

func iterateOverSomething(ctx context.Context, iterations int) {
	for i := 0; i < iterations; i++ {
		if ctx.Err() != nil {
			fmt.Println("Сontext has been canceled, stopping iterateOverSomething...")
			return
		}

		time.Sleep(time.Second * 1)
		fmt.Println("Iteration 1 has been finished", i)
	}
}