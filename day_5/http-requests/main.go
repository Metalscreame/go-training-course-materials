package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//Get()
	MakeRequestPostForm()
}

func MakeRequestPostForm() {
	formData := url.Values{
		"name": {"masnun"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result["form"])
}

func Get() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", body)
}

func Get2() {
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	// 		return slowOperation(ctx)
	req.WithContext(ctx)

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", body)
}
