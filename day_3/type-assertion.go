package main

import "fmt"

/*
Type assertion is a way to get the underlying value an interface holds.
This means if an interface variable is assigned a string then the underlying value it holds is the string.
Here is an example showing how to use type assertion using interfaces.
*/

type B struct {
	s string
}

func main() {
	var i interface{} = B{"a sample string"}

	value, ok := i.(B)
	if !ok {
		fmt.Println("not ok")
	}

	fmt.Println(value.s)




	checkType(i)   // B
	checkType(value.s)

}



func checkType(i interface{}) {
	switch i.(type) {          // the switch uses the type of the interface
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case B:
		fmt.Println("Struct")
	default:
		fmt.Println("Other")
	}
}
