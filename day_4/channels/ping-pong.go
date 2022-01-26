package main

import (
	"fmt"
	"time"
)

/*
Let's write simple ping pong game using channels
 */

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- &Ball{} // game on; toss the ball comment to show deadlock
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
