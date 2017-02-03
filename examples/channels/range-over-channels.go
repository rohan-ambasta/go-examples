package main

import "fmt"

func main() {
	fmt.Println("Range over channels")

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	queue <- "five"
	close(queue)

	for item := range queue {
		fmt.Println(item)
	}
}
