package main

import (
	"fmt"
	"sort"
)

/*
All algorithms in the Go sort package make O(n log n) comparisons in the worst case, where n is the number of elements to be sorted.
Most of the functions are implemented using an optimized version of quicksort.
 */

func main() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]
	fmt.Println(sort.IntsAreSorted(s)) // true



	s = []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// Output: [6 5 4 3 2 1]



	f := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(f)
	fmt.Println(s)


	// Sort with custom comparator
	//Use the function sort.Slice. It sorts a slice using a provided function less(i, j int) bool.
	//To sort the slice while keeping the original order of equal elements, use sort.SliceStable instead.
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]

	/*
	Sort custom data structures
	Use the generic sort.Sort and sort.Stable functions.
	They sort any collection that implements the sort.Interface interface.

	type Interface interface {
	        // Len is the number of elements in the collection.
	        Len() int
	        // Less reports whether the element with
	        // index i should sort before the element with index j.
	        Less(i, j int) bool
	        // Swap swaps the elements with indexes i and j.
	        Swap(i, j int)
	}
	 */

	myFamily := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(myFamily))
	fmt.Println(myFamily) // [{Eve 2} {Alice 23} {Bob 25}]




	/*
	A map is an unordered collection of key-value pairs.
	If you need a stable iteration order, you must maintain a separate data structure.
	This code example uses a slice of keys to sort a map in key order.

	 */

	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// Alice 2
	// Bob 3
	// Cecil 1
}


type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }