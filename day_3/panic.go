package main

import "fmt"

/*

In Go language, panic is just like an exception,
it also arises at runtime.
Or in other words, panic means an unexpected condition
arises in your Go program due to which the execution of your program is terminated.

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds.
To signify events such as end-of-file, Go programs use the built-in error type.
See Error handling best practice and 3 simple ways to create an error for more on errors.

A panic stops the normal execution of a goroutine:

When a program panics, it immediately starts to unwind the call stack.
This continues until the program crashes and prints a stack trace,
or until the built-in recover function is called.
*/

func main() {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("caught panic: ", err)
	//	}
	//}()

	// Creating an array of string type
	myarr := make([]string, 0, 2)

	// Elements are assigned using an index
	myarr[0] = "smth"

	// Accessing the elements of the array using index value
	fmt.Println("Elements of Array:")
	fmt.Println("Element 1: ", myarr[0])

	// Program panics because the size of the array is 3 and we try to access index 5 which is not available in the current slice,
	// So, it gives an runtime error
	fmt.Println("Element 2: ", myarr[5])

	// or you can panic by yourself
	// but you shouldn't do it, this stops execution of the whole program

	panic("decided to panic")

}
