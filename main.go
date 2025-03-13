package main

import (
	"fmt"
	"time"
)

func main() {

	t0 := time.Now()

	channel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		channel <- "data"
	}()
	go func() {
		anotherChannel <- "another-data"
	}()

	select {
	case msg1 := <-channel:
		fmt.Println(msg1)
	case msg2 := <-anotherChannel:
		fmt.Println(msg2)
	}

	printComplete(t0)
}

func printComplete(t time.Time) {
	green := "\033[32m"
	reset := "\033[0m"

	// Print the text in green
	fmt.Printf("\n%scompleted in: %v%s", green, time.Since(t), reset)
}
