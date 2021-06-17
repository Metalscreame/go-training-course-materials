package main

import (
	"fmt"
	"time"
)

/*
Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately
saves the excessive use of memory, acts as a time saver and more importantly, provides better readability of the code.
So basically, a function is a collection of statements that perform some specific task and return the result to the caller.
A function can also perform some specific task without returning anything.
*/

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(multiple(10, 3))

	m := map[string]int{
		"key": 0,
	}

	value, found := getData(m, "key")
	fmt.Println(value, found)

	resString := concat("Hello", "world")
	fmt.Println(resString)

	// function can be a variable. Go is multi-paradign language, so here you may see some ideas from functional programming
	var adder func(x, y int) int
	adder = func(x, y int) int {
		return x + y
	}

	adder(1, 2)

	functionReceiver(adder, 1, 2)

	functionVariable := add
	functionVariable(1, 2)

	fmt.Println(timeNow())
	timeNow = func() time.Time {
		return time.Now().Add(time.Hour * 24)
	}
	fmt.Println(timeNow())
}

// few args
func add(x int, y int) int {
	return x + y // data should be returned if returned data is set. compile error
}

// few args of the same type
func multiple(x, y int) int {
	return x * y
}

// you can return any number of arguments from the function
// but go convention asks us to return maximum 2
func getData(m map[string]int, key string) (int, bool) {
	val, ok := m[key]
	return val, ok
}

// variadic number of args
func concat(stringsToConcat ...string) (result string) { // Can only use '...' as the final argument in the list and it can be only one
	for _, v := range stringsToConcat {
		result += v
	}
	return // will return named value defined at the top of the func
}

func functionReceiver(f func(int, int) int, value1, value2 int) {
	f(value2, value2)
}

var timeNow = func() time.Time {
	return time.Now()
}
