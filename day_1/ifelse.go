package main

import "fmt"

func main() {
	a := true
	if a {
		fmt.Println(a)
	}

	b, c := 1, 2
	if a && b == 1 && c == 2 {
		fmt.Println(a)
	}

	// 1st condition and one of the second behind the brackets
	// for readability you can add new line
	if a &&
		(b == 1 || c == 2) {
		fmt.Println(a)
	}

	mm := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}

	if value, ok := mm["key1"]; ok {
		// value and ok  will be visible only inside of this if block
		fmt.Println(value)
	} else if value == "val2" {
		fmt.Println(value)
	} else {
		fmt.Println("something else")
	}

	// for loops. there is only one for
	//for {  // while true analog. endless loop

	//}

	slice := []int{0, 1, 2, 3, 4, 5}
	count := 0
	for count < 5 {
		if slice[count] < 2 {
			count++
			continue
		}

		fmt.Println(slice[count]) // will skip 0,1
		count++
	}

	// how to iterate over slices
	for i := 0; i < len(slice); i++ {
		if i == 2 {
			break
		}

		fmt.Println(slice[i])
	}

	//  index and !copy! of value under this index
	for idx, value := range slice {
		fmt.Printf("idx %d, val %v\n", idx, value)
	}
	//  if you don't need index
	for _, value := range slice {
		fmt.Println(value)
	}

	//  this is how to iterate over map
	for key, value := range mm {
		fmt.Printf("key - %s, val - %s \n", key, value)
	}

	switch mm["key1"] {
	case "val1":
		fmt.Println("val1")
		break // to move out of switch-case
	case "val2":
		fallthrough // execution flow goes on the lower case. in our case it will be default
	default:
		fmt.Println("nothing is found")
	}

	// equals to
	switch {
	case mm["key1"] == "val1":
		fmt.Println("val1")
		break // to move out of switch-case
	case mm["key1"] == "val2":
		fallthrough // execution flow goes on the lower case. in our case it will be default
	default:
		fmt.Println("nothing is found")
	}


	// break is in the switch and in the for. How to break out of switch inside for loop?
Loop: // goto lable
	for _, val := range mm {
		switch val {
		case "val1":
			fmt.Println("nothing is found")
			break Loop
		}
	}

	// but this is not go-way to use goto
// goto Loop

}
