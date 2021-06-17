package main

import "fmt"

/*
Golang Maps is a collection of UNORDERED pairs of key-value. order can change from compilation to compilation.
It is widely used because it provides fast lookups and values that can retrieve,
update or delete with the help of keys.
It is a reference to a hash table.
 */

func main() {
	// key in map can be any comparable (non-reference type or even a structure that does not contain any reference types
	/*
	comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
	Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.
	 */

	// var m map[string]string not initialized map

	m := make(map[string]int)
	// 	m := make(map[string]map[string]int) you can even create map of maps

	m["key1"] = 2
	fmt.Println(m)

	// to retreive value
	value := m["key1"]
	fmt.Println(value)

	value = m["notExistingKey"] // returns default value
	fmt.Println(value)

	_, exists := m["notExistingKey"] // returns default value
	fmt.Println(exists) // false

	_, exists = m["key1"]
	fmt.Println(exists) // true

	// To delete data from map you need to use
	delete(m, "key1")
	_, exists = m["key1"]
	fmt.Println(exists) // false

	// There is no way to copy map - only using for
}
