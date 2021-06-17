package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var symbolToPrint string
	var widt, length int
	var buffer string
	scanner := bufio.NewScanner(os.Stdin)

	var err error
	fmt.Print("Enter one symbol to print and press Enter: \n")
	scanner.Scan()
	symbolToPrint= scanner.Text()

	if len(symbolToPrint) > 1 {
		os.Exit(1)
	}

	fmt.Printf("Enter the width and press Enter: \n")
	scanner.Scan()
	buffer = scanner.Text()

	//if float check
	widt, err = strconv.Atoi(buffer)
	if err != nil || widt <= 0 {
		os.Exit(1)
	}

	fmt.Printf("Enter the length and press Enter: ")
	scanner.Scan()
	buffer = scanner.Text()

	length, err = strconv.Atoi(buffer)
	if err != nil || length <= 0 {
		os.Exit(1)
	}

	fmt.Printf("\n\n")

	for i := 1; i < widt+1; i++ {
		if i%2 == 0 {
			fmt.Print(" ")
		}
		for j := 0; j < length; j++ {
			fmt.Print(symbolToPrint)
		}
		fmt.Print("\n")
	}
}