package main

import (
	"fmt"
	"time"
)

func player(name string, table chan string) {
	for {
		ball, ok := <-table
		if !ok {
			return
		}
		fmt.Printf("%s got the ball: %s\n", name, ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball // send the ball back
	}
}

func main() {
	table := make(chan string) // create a channel to pass the ball

	// initiate 2 players using goroutine
	go player("Dani", table)
	go player("Budi", table)

	// start the game
	table <- "ping"

	time.Sleep(1 * time.Second) // the game will last for 1 second
	close(table)                // close the channel to end the game
}
