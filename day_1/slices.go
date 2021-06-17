package main

import "fmt"

/*
In Go language slice is more powerful, flexible, convenient
than an array, and is a lightweight data structure.
Slice is a variable-length sequence which stores
elements of a similar type,
you are not allowed to store different type of elements
in the same slice.
*/

func main() {
	// var a [5]int - array
	var emptySlice []int
	fmt.Println(emptySlice)
	fmt.Println(len(emptySlice))
	// introducing function cap - capacity shows
	//maximum capacity of current slice
	// under the hood slice is the same array.
	fmt.Println(cap(emptySlice))

	var notEmptySlice []int
	// notEmptySlice[0] = 1 // runtime error: index out of range [0] with length 0

	// make function allocates memory in heap. until that slice is not initialized
	// why do we need cap? Let me show you how to initialize a slice first.
	notEmptySlice = make([]int, 0, 6) // this tells the compiler - initialize a slice with 0 len and 6 max capacity
	fmt.Println(cap(notEmptySlice))    // 6

	notEmptySlice = append(notEmptySlice, 0, 1, 2, 3, 4)
	pointToNotEmpty := notEmptySlice // this is not copy of an slice. it copies the pointer to the initial slice
	fmt.Println(notEmptySlice)
	fmt.Println(pointToNotEmpty)

	fmt.Println(cap(notEmptySlice)) // 6
	fmt.Printf("notEmptySlice address %p\n", notEmptySlice)
	fmt.Printf("pointToNotEmpty address %p\n", pointToNotEmpty)

	// if you change the value in notEmptySlice - it will be changed in pointToNotEmpty
	notEmptySlice = append(notEmptySlice, 5) // this is one of the slices gotchas read about it in the link that will be provided
	notEmptySlice[1] = 999                   // will do the change in both

	fmt.Println(notEmptySlice)
	fmt.Println(pointToNotEmpty)
	fmt.Println(cap(notEmptySlice)) // 6

	// when slice reach it's capacity append function under the hood alocates new x2 sized slice, moves all the  and returns it.
	// this creates completely new slice and data from old one is copied into the new one
	notEmptySlice = append(notEmptySlice, 6)
	fmt.Println(cap(notEmptySlice)) // 10
	fmt.Printf("notEmptySlice now has new address %p\n", notEmptySlice)
	fmt.Printf("pointToNotEmpty points to old address %p\n", pointToNotEmpty)

	fmt.Println(notEmptySlice)
	fmt.Println(pointToNotEmpty)
	notEmptySlice[0] = 111111
	fmt.Println(notEmptySlice)
	fmt.Println(pointToNotEmpty)
	// 	pointToNotEmpty := notEmptySlice

	var copyOfSlice = make([]int, len(notEmptySlice), cap(notEmptySlice)) // if you put an empty slice there - copy won't work
	copy(copyOfSlice, notEmptySlice)                                      // copies into
	fmt.Printf("notEmptySlice points to address %p\n", notEmptySlice)
	fmt.Printf("copyOfSlice points to address %p\n", copyOfSlice)
	fmt.Println(notEmptySlice)
	fmt.Println(copyOfSlice)

	notEmptySlice = append(notEmptySlice, copyOfSlice...) // this unpacks values from this slice into
	fmt.Println(notEmptySlice)

}
