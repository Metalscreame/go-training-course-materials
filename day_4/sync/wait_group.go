package main

import "sync"


/*
sync.WaitGroup tends also to be used quite frequently.
It’s the idiomatic way for a goroutine to wait for the completion of a collection of goroutines.
sync.WaitGroup holds an internal counter.
If this counter is equal to 0, the Wait() method returns immediately.
Otherwise,it is blocked until the counter is 0.
To increment the counter we have to use Add(int).
To decrement it we can either use Done() (that will decrement by 1) or
the same Add(int) method with a negative value.
In the following example, we will spin up eight goroutines and wait for their completion:
 */
func main() {

	wg := &sync.WaitGroup{}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			// Do something
			wg.Done()
		}()
	}

	wg.Wait()
}