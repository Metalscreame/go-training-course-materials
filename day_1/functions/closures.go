package main

import "fmt"

/*
Go supports anonymous functions, which can form closures.

A closure is a special type of anonymous function that
references variables declared outside of the function itself. It is similar to accessing global variables which are available before the declaration of the function.

Anonymous functions are useful when you want to define a function inline without having to name it.

We call intSeq, assigning the result (a function) to nextInt.
This function value captures its own i value, which will be updated each time we call nextInt.
 */

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}