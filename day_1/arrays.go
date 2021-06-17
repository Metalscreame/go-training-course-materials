package main

import "fmt"

/*
from now on i will start using package fmt to print data
let's just skip the explanation for now

In Go, an array is a numbered sequence of elements of a specific and fixed length.
They are rarely used but you should now about them to better understand how slices, that are used instead of arrays
works.
It's low level language type

*/
func main() {
	// fixed size
	// it's filled with default values of ints 0 0 0 0 0
	var a [5]int
	fmt.Println(a)

	//len of array is part of it's type, so if function waits for a 3 len array, you won't be able to put 4 len array there

	fmt.Println(len(a))
	// myFunc(a) // compile error  cannot use a (type [5]int) as type [4]int in argument to myFunc

	// if you don't want to count compiler will count this for you
	// on compilation
	a2 := [...]int{1, 2, 3}
	fmt.Println(a2)
	fmt.Println(len(a2))


	var aa [3][3]int
	fmt.Println(aa)
	fmt.Println(len(aa))

	aa[2][1] = 20
	fmt.Println(aa)
	fmt.Println(aa[2][1])
}

func myFunc(ar [4]int)  {

}
