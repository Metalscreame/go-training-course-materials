package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)


	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
