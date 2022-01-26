package main

import (
	"fmt"
	"os"
)

// A defer statement defers the execution of a
//function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately,
//but the function call is not executed until the
//surrounding function returns.

func main() {
	for i := 0; i < 10; i++ {
		// defers are executed in last-in-first-out order
		defer fmt.Println(i)
	}

	var pointer int
	pointer = 0

	defer func(numPoint *int) {
		fmt.Println("number in pointer", *numPoint)
	}(&pointer)

	// this will change number in defered call
	pointer = 999

	fmt.Println("Hello")


	// best real world example
	file, err := os.Open("closures.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()


}
