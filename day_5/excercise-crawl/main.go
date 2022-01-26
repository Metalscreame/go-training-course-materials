package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const stringToSearch = "concurrency"

var sites = []string{
	"https://google.com",
	"https://itc.ua/",
	"https://twitter.com/concurrencyinc",
	"https://twitter.com/",
	"http://localhost:8000",
	"https://github.com/bradtraversy/go_restapi/blob/master/main.go",
	"https://www.youtube.com/",
	"https://postman-echo.com/get",
	"https://en.wikipedia.org/wiki/Concurrency_(computer_science)#:~:text=In%20computer%20science%2C%20concurrency%20is,without%20affecting%20the%20final%20outcome.",
}

type SiteData struct {
	data []byte
	uri  string
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	resultsCh := make(chan *SiteData, len(sites))

	for _, site := range sites {
		go RequestWorker(ctx, resultsCh, site)
	}

	go Searcher(ctx, cancel, resultsCh)

	select {case <-ctx.Done():}
}

func Searcher(ctx context.Context, cancelFunc context.CancelFunc, res chan *SiteData) {
	for {
		select {
		case data := <-res:
			go func(res *SiteData) {
				if strings.Contains(string(res.data), stringToSearch) {
					fmt.Printf("'%v' string is found in %s\n", stringToSearch, res.uri)
					cancelFunc()
					return
				}

				fmt.Printf("Nothing found in %s\n", res.uri)
			}(data)
		case <-ctx.Done():
			fmt.Println("exiting from searcher...")
			return
		}
	}
}

func RequestWorker(ctx context.Context, outChan chan *SiteData, uri string) {
	fmt.Printf("starting sending request to %v\n", uri)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
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
	outChan <- &SiteData{
		data: body,
		uri:  uri,
	}
}
