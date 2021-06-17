// Golang program to compare equality
// of struct, slice, and map
package main

import (
	"fmt"
	"reflect"
)

type compared struct {
	a int
}

type structeq struct {
	X int
	Y string
	Z []int
}

func main() {
	a1 := compared{
		a: 1,
	}
	a2 := compared{
		a: 2,
	}
	fmt.Println(a1 == a2)

	s1 := structeq{X: 50,
		Y: "Y",
		Z: []int{1, 2, 3},
	}
	s2 := structeq{X: 50,
		Y: "Y",
		Z: []int{1, 2, 3},
	}

	// comparing struct
	if reflect.DeepEqual(s1, s2) {
		fmt.Println("Struct is equal")
	} else {
		fmt.Println("Struct is not equal")
	}

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3, 4}

	// comparing slice
	if reflect.DeepEqual(slice1, slice2) {
		fmt.Println("Slice is equal")
	} else {
		fmt.Println("Slice is not equal")
	}

	map1 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}
	map2 := map[string]int{
		"x": 10,
		"y": 20,
		"z": 30,
	}

	// comparing map
	if reflect.DeepEqual(map1, map2) {
		fmt.Println("Map is equal")
	} else {
		fmt.Println("Map is not equal")
	}
}
