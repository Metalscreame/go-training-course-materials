package main

import (
	"encoding/json"
	"fmt"
)

/*
At times you want to read a JSON file and extract just a few fields.

If you want to avoid defining a struct, you could decode it into a map of generic types and cast them appropriately.
*/

func main() {
	b := []byte(`{"earnings":1000.50,"names":["Joe","Jane"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	earnings := dat["earnings"].(float64)
	fmt.Println(earnings)
}