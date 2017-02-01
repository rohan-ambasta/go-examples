package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		if i > 1 {
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println(from, ":", i)
	}
}

func main() {
	fmt.Println("GoRoutines example")

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
