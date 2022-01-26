package main

import "fmt"

/*
Pointers in Go programming language or Golang is a
variable which is
used to store the memory address of another variable.
*/

func main() {

	i := 5
	b := &i // in b now stored memory address of i variable
	fmt.Println("address ", b)
	fmt.Println("value", i)

	*b = 6 // this changes the value under this memory
	fmt.Println("address ", b)
	fmt.Println("value", i)

	// everything in go is send by value to a functions
	// but some types (slice, map, array) has a pointer to a memory inside
	// technically speaking when you send a slice to a function - you're sending pointer to a slice
	slice := []int{1, 3, 4}
	changer(slice)
	fmt.Println(slice)

	ss := structPoint{
		i:     0,
		slice: []int{1, 3, 4},
	}
	changerStruct(ss)
	fmt.Println(ss)

	/*
	func new ¶
	func new(Type) *Type
	The new built-in function allocates memory.
	The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
	 */

	pntr:= new(structPoint)
	fmt.Println(pntr)
}

type structPoint struct {
	i     int
	slice []int
}

func changer(items []int) {
	items[0] = 999
}
func changerStruct(s structPoint) {
	s.slice[0] = 999 // will be changed
	s.i = 999        // will not
}


/*
nil is a predeclared identifier in Go that represents zero values
for pointers, interfaces, channels, maps, slices and function
types. nil being the zero value of pointers and interfaces,
uninitialized pointers and interfaces will be nil.
 */

type A struct {
	str string
}
func main() {
	s := getA()
	fmt.Println(s.str) // This will crash
}
func getA() *A {
	return nil // This WILL compile
}

func main() {
	var myPointer *int
	myNum := 2

	myPointer = &myNum
	pointerTest(myPointer)
}

func pointerTest(p *int)  {
	fmt.Println(p == nil)
}
