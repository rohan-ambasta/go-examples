package main

import "fmt"

func main() {
	//Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those
	// values into another goroutine.
	fmt.Println("Channels example")

	//Create a new channel with make(chan val-type).
	//Channels are typed by the values they convey.
	messages := make(chan string)

	//add a string value in messages channel
	go func() { messages <- "ping" }()

	//receive the string value from messages channel
	//By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message
	// without having to use any other synchronization.
	msg := <-messages
	fmt.Println(msg)
}
