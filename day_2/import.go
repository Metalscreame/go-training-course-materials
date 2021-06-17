package main

import (
	"fmt"

	_ "github.com/Metalscreame/go-training/day_2/import_pack"
	// . "github.com/Metalscreame/go-training/day_2/import_pack"

)

/*

Variables may also be initialized using init functions.

func init() { … }
Multiple such functions may be defined. They cannot be called from inside a program.

A package with no imports is initialized
by assigning initial values to all its package-level variables,
followed by calling all init functions in the order they appear in the source.
Imported packages are initialized before the package itself.
Each package is initialized once, regardless if it’s imported by multiple other packages.
It follows that there can be no cyclic dependencies.

Package initialization happens in a single goroutine, sequentially, one package at a time.


The order that the respective filenames were passed to the Go compiler.

The Go spec says "build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler"
so it's a safe bet that go build does exactly that, and the inits will run in A-B-C order
 */

func main() {
	//fmt.Println(PublicVar)
	fmt.Println("Hi")
}
