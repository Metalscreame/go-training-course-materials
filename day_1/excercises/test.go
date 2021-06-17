package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter string")

	var data string
	fmt.Scan(&data)

	arr := strings.Split(data, ",")

	evenCount := 0
	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			break
		}

		if num%2 == 0 {
			evenCount++
		}
	}
	fmt.Println("Number of even numbers is: ", evenCount)
}
