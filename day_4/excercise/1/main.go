//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"strings"
	"time"
)

func producer(stream Stream, dataCh chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(dataCh)
			return
		}

		dataCh <- tweet
	}
}

func consumer(dataCh chan *Tweet, done chan struct{}) {
	for t := range dataCh {
		fmt.Println(t)
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
			continue
		}

		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
	done <- struct{}{}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	done := make(chan struct{})
	dataCh := make(chan *Tweet, len(stream.tweets))

	// Producer
	go producer(stream, dataCh)

	// Consumer
	go consumer(dataCh, done)

	<-done
	fmt.Printf("Process took %s\n", time.Since(start))
}

// Tweet defines the simlified representation of a tweet
type Tweet struct {
	Username string
	Text     string
}

// IsTalkingAboutGo is a mock process which pretend to be a sophisticated procedure to analyse whether tweet is talking about go or not
func (t *Tweet) IsTalkingAboutGo() bool {
	// simulate delay
	time.Sleep(330 * time.Millisecond)

	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher
}
