package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type myStruct struct {
	Name string
}

/*
In Go it’s idiomatic to communicate
errors via an explicit, separate return value.
This contrasts with the exceptions used in languages
like Java and Ruby and the overloaded single result /
error value sometimes used in C.
Go’s approach makes it easy to see which functions return
errors and to handle them using the same language constructs
employed for any other, non-error tasks.

By convention, errors are the last return value and have type error, a built-in interface.
A nil value in the error position indicates that there was no error.
*/
func main() {
	m := myStruct{
		Name: "name",
	}

	bytes, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))

	err = json.Unmarshal([]byte(`{wrong}`), &m)
	if err != nil {
		fmt.Println(err.Error())
	}

	// you can define your own errors and then later make this checks
	err = finder()
	if  err == ErrNotFound {
		fmt.Println(err.Error())
	}

	// The errors.Is function compares an error to a value.
	if errors.Is(err, ErrNotFound) {
		// something wasn't found
	}

	/*
	Wrapping an error with %w makes it available to errors.Is:

	err := fmt.Errorf("access denied: %w", ErrPermission)
	...
	if errors.Is(err, ErrPermission) ...

	There is another method errors.As - i will send a link to the
	article for you to read about it
	 */

	// error type is basically the type that just implements Error interface
	// everything that implements Error interface
	c := CustomError{customString: "go course"}
	// err can be compared to an struct that implements error method
	err = c
	if err == c {
		fmt.Println(err)
	}
}





func finder() error {
	// doing something and error happened
	return ErrNotFound
}

var ErrNotFound = errors.New("not found")

type CustomError struct {
	customString string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("custom error %s", c.customString)
}


