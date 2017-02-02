package main

import (
	"fmt"
	"time"
)

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine
// that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("\nWorker:\nWorking...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	fmt.Println("Channel Synchronization example")
	done := make(chan bool)

	// Start a worker goroutine, giving it the channel to notify on.
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.
	if <-done {
		fmt.Println("\nMain:\nme too !")
	}
}
