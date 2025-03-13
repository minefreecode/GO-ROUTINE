package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func main() {
	t0 := time.Now()

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)
	fmt.Println(<-done)

	printComplete(t0)
}

func printComplete(t time.Time) {
	green := "\033[32m"
	reset := "\033[0m"

	// Print the text in green
	fmt.Printf("\n%scompleted in: %v%s", green, time.Since(t), reset)
}
