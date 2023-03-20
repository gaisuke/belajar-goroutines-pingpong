package main

import (
	"fmt"
	"math/rand"
	"time"
)

func player(name string, table chan string, ref chan string) {
	for {
		ball, ok := <-table
		if !ok {
			return
		}
		fmt.Printf("%s got the ball: %s\n", name, ball)
		time.Sleep(100 * time.Millisecond)

		if rand.Intn(100)%13 == 0 {
			fmt.Printf("%s dropped the ball\n", name)
			ref <- name // send the name of the player who dropped the ball
			return
		}

		// alternate the ball and send it back to the table
		if ball == "ping" {
			table <- "pong"
		} else {
			table <- "ping"
		}
	}
}

func main() {
	table := make(chan string) // create a channel to pass the ball
	ref := make(chan string)   // create a channel to receive the name of the player who dropped the ball

	// initiate 2 players using goroutine
	go player("Dani", table, ref)
	go player("Budi", table, ref)

	// start the game
	table <- "ping"

	time.Sleep(20 * time.Second) // the game will last for 20 second
	close(table)                 // close the table to end the game
	close(ref)                   // close the ref to end the game
}
