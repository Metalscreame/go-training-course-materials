package main

func main() {

	// ways to declare variables
	var i int = 10 // 32 or 64 bit dependent on your platform
	var i32 int32 = 10
	var i64 int64 = 10 // int8, int 16
	var autoInt = -10
	var unasignedBigInt uint64 = 1<<64 - 1 // bigint

	/*
	go has one interesting feature. usefull, sometimes annoying until you get used to it
	every variable should be used - if no - can't compile
	it helps with understanding what's going on in your program
	you will always know what that variable is used
	*/
	println("integers", i, i32, i64, autoInt, unasignedBigInt)

	var p float32 = 3.1457 // only float32 float64
	println("floats ", p)

	// booleans
	var b bool = true
	println("floats ", b)

	// strings
	// string holds bytes
	var iAmSting string = "string\n\t"
	var rawString = `
					{
						"json": "you can put every symbol''%&#*!!#/n/t"
					}
					`
	println(iAmSting, rawString)

	// binary data
	var rawBinary byte = '\x99'
	println("raw binary", rawBinary)

	// rune represent unicode codepoints. For example, the rune literal 'a' is actually the number 97.
	var character rune = 'a'
	println("rune ", character) // prints 97

	// short declarations
	myInt := 2
	// myInt := 2 - compile error can't declare variable with same name
	myInt = 10 // reassigning

	// variables declaration block
	var (
		myInt2 = 3
		myVar  = 2 + 3i // complex numbers, mathematics..
	)

	println(myInt, myInt2, myVar)

	// default values
	var (
		ii int     // 0
		f  float32 // +0.000000e+000
		s  string  // ""
		bb bool    // false
	)

	println(ii, f, s, bb)

	_ = 2

	// multiple declaration
	var a, c, d int = 0, 2, 3

	println(a, c, d)

	firstNameVar, secondNameVar := "Roman", "MySecondName"
	println(firstNameVar, secondNameVar)

	// to swap variables use
	firstNameVar, secondNameVar = secondNameVar, firstNameVar
	println(firstNameVar, secondNameVar)

	// you can also use expressions while declarating vars
	sum := 2 + 1 + 2 + 3
	println(sum)
}
