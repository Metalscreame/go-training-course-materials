package main

import "fmt"

type example struct {
	flag                           bool
	PublicField, randomStringField string
	slice                          []int
}


/*
A structure or struct in Golang is a user-defined type
that allows to group/combine items of possibly different types into a single type.
Any real-world entity which has some set of properties/fields can be represented as a struct.
This concept is generally compared with the classes in object-oriented programming.
 */
func main() {
	// declare struct with default initial values
	var e example
	//e := example{}

	e = example{
		flag:        true,
		PublicField: "data",
		slice:       []int{1, 3, 3},
	}

	fmt.Println(e)
}
