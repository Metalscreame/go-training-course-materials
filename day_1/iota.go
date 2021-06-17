package main

// autoincrement
const (
	// but if iota is second var it will start
	//count from the place it first appeared, which is 1
	zero = iota // first iota always 0
	one  = iota
	two
	_
	four // 4
)

const (
	_         = iota // skip first
	KB uint64 = 1 << (10 * iota) // bytes in KB .. etc
	MB
	GB
	TB
)



func main() {
	println(zero)
	println(one)
	println(two)


	println(KB)
	println(MB)
	println(GB)
	println(TB)
}
