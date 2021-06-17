package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c Circle) circleArea() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) updateRadius(r float64) {
	c.r = r
}

func main() {
	c := Circle{x: 1, y: 2, r: 5}

	area := c.circleArea()
	fmt.Println(area)

	c.updateRadius(20)

	area = c.circleArea()
	fmt.Println(area)
}
/*
CacheValue receiver makes a copy of the type and pass it to the function. The function stack now holds an equal object but at a different location on memory.
That means any changes done on the passed object will remain local to the method. The original object will remain unchanged.
Pointer receiver passes the address of a type to the function. The function stack has a reference to the original object.
So any modifications on the passed object will modify the original object.

The Pointer receiver avoids copying the value on each method call.
This can be more efficient if the receiver is a large struct,
CacheValue receivers are concurrency safe, while pointer receivers are not concurrency safe.
Hence a programmer needs to take care of it.
 */