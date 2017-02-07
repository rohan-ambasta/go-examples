package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ticker Example")

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for tick := range ticker.C {
			fmt.Println("Tick at ", tick)
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
