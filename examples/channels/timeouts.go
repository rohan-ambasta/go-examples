package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timouts using channels and select")

	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println("received ", res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println("received ", res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
