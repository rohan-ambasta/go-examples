package main

import "fmt"

func main() {
	fmt.Println("Channel Buffering example")

	// create a string channel with buffer size = 2
	messages := make(chan string, 2)
	messages <- "msg1"
	messages <- "msg2"
	// we get a deadlock if we exceed the buffer size
	//messages <- "msg3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fetching the third value will display a deadlock error
	// because there is no goroutine with a send channel
	//fmt.Println(<- messages)

	// create an un-buffered channel
	//messages1 := make(chan string)
	// we will get a deadlock with the below code because there is no
	// go routine with a receive channel
	//messages1 <- "msg3"
}
