package main

import "fmt"

func main() {
	fmt.Println("Channels example")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
