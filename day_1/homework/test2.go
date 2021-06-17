package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Context struct {
	Min int
	Max int
}

type Result struct {
	EasyFormula int
	HardFormula int
}

const (
	errorSigned     = "Numbers cant be signed, or float or 0"
	errorSameLength = "Must be the same lengths"
)

func main() {
	var context Context

	fmt.Print("Enter context, example:  {\"min\":320123,\"max\":320320}\n")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	err := json.Unmarshal([]byte(line), &context)
	simpleErrorChecker(err, "")

	if context.Max <= 0 || context.Min <= 0 {
		errHandler(errors.New(errorSigned))
		os.Exit(1)
	}

	if len(strconv.Itoa(context.Max)) < 6 || len(strconv.Itoa(context.Min)) < 6 {
		errHandler(errors.New("Cant have less than 6 numbers"))
		os.Exit(1)
	}

	if context.Min > context.Max {
		errHandler(errors.New("Min cant be higher than max"))
		os.Exit(1)
	}

	//same length check
	if len(strconv.Itoa(context.Max)) != len(strconv.Itoa(context.Min)) {
		errHandler(errors.New(errorSameLength))
		os.Exit(1)
	}

	var result Result
	result.EasyFormula = easyWay(context.Min, context.Max)
	result.HardFormula = hardWay(context.Min, context.Max)

	r, _ := json.Marshal(result)
	fmt.Printf("%s", r)

}

func easyWay(min, max int) (count int) {
	var isLucky bool

	for i := min; i <= max; i++ {
		isLucky = false
		digits := make([]int, 0)

		//getting digits
		currentNumber := i
		for currentNumber > 0 {
			digit := currentNumber % 10
			digits = append(digits, digit)
			currentNumber /= 10
		}

		var firstThreeSum, secondThreeSum int

		for j := 0; j < len(digits)/2; j++ {
			firstThreeSum += digits[j]
		}

		if len(digits)%2 == 0 {
			for j := len(digits) / 2; j < len(digits); j++ {
				secondThreeSum += digits[j]
			}
		} else { // if length of a ticket is not even
			for j := (len(digits) / 2) + 1; j < len(digits); j++ {
				secondThreeSum += digits[j]
			}
		}

		if firstThreeSum == secondThreeSum {
			isLucky = true
		}
		if isLucky {
			count++
		}
	}
	return count
}

func hardWay(min, max int) (count int) {
	var isLucky bool

	for i := min; i <= max; i++ {
		isLucky = false
		digits := make([]int, 0)
		var evenSum, oddSum int

		//getting digits
		currentNumber := i
		for currentNumber > 0 {
			digit := currentNumber % 10
			digits = append(digits, digit)
			currentNumber /= 10
		}
		for _, element := range digits {
			if element%2 == 0 {
				evenSum += element
			} else {
				oddSum += element
			}
		}
		if evenSum == oddSum {
			isLucky = true
		}
		if isLucky {
			count++
		}
	}
	return count
}

type ErrorResponse struct {
	Status string
	Reason string
}

func simpleErrorChecker(err error, possibleMessage string) {
	if err != nil && possibleMessage != "" {
		errHandler(errors.New(possibleMessage))
		os.Exit(1)
	} else if err != nil {
		errHandler(err)
		os.Exit(1)
	}
}

func errHandler(err error) {
	errResponse := &ErrorResponse{Status: "failed", Reason: err.Error()}
	result, _ := json.Marshal(errResponse)
	fmt.Println(string(result))
}
