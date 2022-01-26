package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name    int      `json:"name"`
	Entries []string `json:"entries"`
}

func main() {

	str := `{"name": 1, "entries": ["listObj1", "listObj2"]}`

	res := Response{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		fmt.Printf("error %v", err)
		return
	}

	fmt.Println(res)
}
