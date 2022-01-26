package main

import (
	"fmt"
	"os"
)

/*
When to use parameters
Definitely the first form is preferred if you plan to change the value
of the variable which you don't want to observe in the function.

This is the typical case when the anonymous functionis inside a for
loop and you intend to use the loop's variables, for example:
*/

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

/*
Without passing the variable i you might observe printing 10 ten times.
With passing i, you will observe numbers printed from 0 to 9.

When not to use parameters
If you don't want to change the value of the variable,
it is cheaper not to pass it and thus not create another copy of it.
This is especially true for large structs. Although if you later alter the code and modify
the variable, you may easily forget to check its effect on the closure and get unexpected
results.

Also there might be cases when you do want to observe changes
made to "outer" variables, such as:
 */

func GetRes(name string) (*os.File, error) {
	res, err := os.Open(name)
	if err != nil {
		return res, err
	}

	closeres := true
	defer func() {
		if closeres {
			res.Close()
		}
	}()

	// Do other stuff

	// Everything went well, return res, but
	// res must not be closed, it will be the responsibility of the caller
	closeres = false
	return res, nil // res will not be closed
}