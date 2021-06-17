package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter string")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	arr := strings.Split(scanner.Text(), " ")

	if len(arr) != 4 {
		fmt.Println("data has not enough numbers, len", len(arr))
		return
	}

	var res []string
	for i, v := range arr {
		if len(v) != 4 {
			fmt.Println("data has wrong len ", len(v))
			return
		}

		_, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			break
		}

		if i != len(arr)-1 {
			res = append(res, "****")
			continue
		}

		res = append(res, v)
	}

	fmt.Println(strings.Join(res, " "))
}
