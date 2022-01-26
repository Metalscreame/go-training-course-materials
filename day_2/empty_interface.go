package main

import "fmt"

/*
When an interface has zero methods, it is called
an empty interface. This is represented by interface{}.
Since the empty interface has zero methods, all types implement this interface implicitly.
Have you wondered how does the Println function from the fmt built-in package accepts the different types of value as arguments?
This is possible because of an empty interface. Let’s see the signature of Println function
 */

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
