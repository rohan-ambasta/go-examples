package main

import "fmt"

// here pings is receive only channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// here pings is a send only channel
// pongs is a receive only channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg // can also be written as pongs<- <-pings instead of using msg variable
}
func main() {
	fmt.Println("Channel directions example")
	pings := make(chan string)
	pongs := make(chan string)
	go ping(pings, "hello")
	go pong(pings, pongs)
	fmt.Println(<-pongs)

}
