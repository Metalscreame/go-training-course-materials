package main

import (
	"fmt"
	"sync"
)
/*
sync.Map is a concurrent version of Go map where we can:
Add elements with Store(interface{}, interface{})
Retrieve elements with Load(interface) interface{}
Remove elements with Delete(interface{})
Retrieve or add an element if it did not exist before with
LoadOrStore(interface{}, interface{}) (interface, bool).
The returned bool is true if the key was present in the map before.
Iterate on the elements with Range

https://habrastorage.org/web/fc8/179/905/fc81799050554f20a7b28314b68afe47.png
 */
func main() {
	m := &sync.Map{}

	// Put elements
	m.Store(1, "one")
	m.Store(2, "two")

	// Get element 1
	value, contains := m.Load(1)
	if contains {
		fmt.Printf("%s\n", value.(string))
	}

	// Returns the existing value if present, otherwise stores it
	value, loaded := m.LoadOrStore(3, "three")
	if !loaded {
		fmt.Printf("%s\n", value.(string))
	}

	// Delete element 3
	m.Delete(3)

	// Iterate over all the elements
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %s\n", key.(int), value.(string))
		return true
	})
}