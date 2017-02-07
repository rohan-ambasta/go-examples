package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Using channels with select")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//make this interesting and observe
	//time.Sleep(time.Second * 5)

	for i := 0; i < 1; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		default:
			fmt.Println("default msg")
		}
	}
}
