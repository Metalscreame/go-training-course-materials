package main

import (
	"context"
	"fmt"
	"time"
)

/*
Contexts with timeout are mostly used when you want to make an external
request, such as querying on your database or communicating with another
service (either with HTTP, gRPC, or any other protocol).
Okay! Let me show you some codes examples using contexts
and demonstrate where you may use any of them in real world.
 */

func main() {
	// Channel used to receive the result from doSomething
	dataCh := make(chan string)

	// Creates a new context with timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go doSomething(ctx, 7, dataCh)

	select {
	// I'm waiting to check if the context has timed out
	case <-ctx.Done():
		fmt.Printf("Context cancelled: %v\n", ctx.Err())
	// I'm waiting to check if I receive the results from doSomething
	case result := <-dataCh:
		fmt.Printf("processing finished: %v\n", result)
	}
}

func doSomething(ctx context.Context, timeSleep time.Duration, ch chan string) {
	fmt.Println("Starting processing something heavy...")
	time.Sleep(time.Second * timeSleep)
	fmt.Println("Waiking up...")
	ch <- "Did something!"
}
