package main

import (
	"fmt"

	// this imports imports data into this file, not the whole package
	"github.com/Metalscreame/go-training/day_1/exampler"
)

/*
A package is the smallest unit of private encapsulation in Go.

All identifiers defined within a package are visible throughout
that package.
When importing a package you can access only its exported identifiers.
An identifier is exported if it begins with a capital letter.

Exported and unexported identifiers are used to describe the public interface of a package and to guard against certain programming errors.
*/

func main() {
	// you can import only stuff that starts with capital letter
	imported := exampler.GlobalPublicVariable
	fmt.Println(imported)

	fmt.Println(packageLevelVar)

	otherFunc(packageLevelVar)
}

var packageLevelVar = "packageLevel"

func otherFunc(data string) {
	fmt.Println(packageLevelVar)
	fmt.Println("Got from func ", data)

	if data != "" {
		data = "myChangedData"
		createdVar := "" // visible only in this if block
		data += createdVar
	}
	fmt.Println("data is changed to ", data)

	if data != "" {
		// shadowing
		data := "completelyDifferent"
		createdVar := ""
		data += createdVar
	}
	fmt.Println("data is unchanged ", data)

}
