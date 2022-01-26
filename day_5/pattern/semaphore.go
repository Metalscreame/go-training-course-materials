package main

import (
	"fmt"
	"sync"
	"time"
)

// remoteDeleteEmployeeRPC will delete an employee over the network.
func remoteDeleteEmployeeRPC(num int) {
	fmt.Println(num)
	time.Sleep(time.Millisecond * 300)
}

// sem is a channel that will allow up to n concurrent operations.
// if we use 2 it will be alsmost as if we're iterating over slice
// if 10 - output will change
var sem = make(chan struct{}, 10)

func main() {
	var ints = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	for _, num := range ints {
		sem <- struct{}{}
		wg.Add(1)

		go func(num int) {
			defer func() {
				wg.Done()
				<-sem
			}()

			remoteDeleteEmployeeRPC(num)
		}(num)
	}
	wg.Wait()
}
