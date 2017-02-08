package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future -
// tickers are for when you want to do something repeatedly at regular
// intervals. Here’s an example of a ticker that ticks periodically until we stop it.
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
