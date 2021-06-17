package main

// const are initalized on startup and can't be
//changed on runtime

const myConstIntNotTyped = 123
const piNumber = 3.14
const integerConst int32 = 2

const (
	myConstString = "123"
	myChar        = '1'
)


func main() {
	println(myConstIntNotTyped + piNumber)

	// can't do that as constant is typed, has to cast
	// println(piNumber+integerConst)
	println(piNumber + float64(integerConst))
}
