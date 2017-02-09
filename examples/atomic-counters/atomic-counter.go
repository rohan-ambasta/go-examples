package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Atomic counters")

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	opsValue := atomic.LoadUint64(&ops)
	fmt.Println(opsValue)
}
