package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)
/*
In Go, a string is in effect a read-only slice of bytes.
And also, Go source code is UTF-8,
so the source code for the string literal is UTF-8 text.
I will send a link to the good article related to unicode strings in go
All you need to know about strings to be ok i will explain now
 */

func main() {
	str := "im a sting"
	// str[1] = "2" compile error
	str += " addition"
	str = fmt.Sprintf("formatted '%s'", str) // will print
	fmt.Println(str)

	// to replace something use
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))

	// if you want to repace char at specific index = you cannot do that easey, you'll need to write somth like this
	replace := func (str string, replacement rune, index int) string {
		return str[:index] + string(replacement) + str[index+1:]
	}

	beforeString:= "before"
	after := replace(beforeString, 'a', 0)
	fmt.Println(after)

	// unicode
	chString := "福寿"           // 2 symbols
	fmt.Println(len(chString)) // len - 6. Takes 6 bytes to store those symbols
	// so if you'll have strings out of utf-8 table you'll need to be avare of this behaviour

	// this will return numbers of characters no matter what encoding is used
	fmt.Println(utf8.RuneCountInString(chString)) // 2

	var character rune
	character = 'a'
	fmt.Println(character) // will print unicode representation of this char
	fmt.Println(string(character))

	// strings are unchangable type
	// chString[0] = "2"
	chString = chString + "me" // but you can concatenate stgings
	chString += "e"

	// you can iterate only by bytes. using string
	for idx, value := range chString {
		fmt.Printf("%#U at position %d\n", value, idx)
	}

	fmt.Println(chString[1]) // byte value

	// if you want to iterate over bytes - cast to byte
	b := []byte(chString)
	for idx, value := range b {
		fmt.Printf("%#U at position %d\n", value, idx)
	}

	// if you want to iterate over
	r := []rune(chString)
	for idx, value := range r {
		fmt.Printf("%#U at position %d\n", value, idx)
	}



	// you can't convert int to string and string to int to get correct values
	// so to convert int -> string and backward
	number, _ := strconv.Atoi("2")

	fmt.Println(number)
	numberString := strconv.Itoa(123)
	fmt.Println(numberString)

	// strings package overview

	//the most efficient comparing way
	fmt.Println(strings.EqualFold("*", "*"))

	indx := strings.Index("1235", "2")
	fmt.Println(indx)

	var strBuilder strings.Builder
	strBuilder.WriteString("ignition\n")
	strBuilder.WriteString("ignition 2")
	fmt.Println(strBuilder.String())
}
