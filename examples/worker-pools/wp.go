package main

import (
	"fmt"
	"time"
)

func test1(done chan bool, items chan int, testId int) {
	for {
		//more is used because we are closing items channel
		item, more := <-items
		if more {
			fmt.Println("take item ", testId, item)
			time.Sleep(time.Second)
			fmt.Println("remove item ", testId, item)
		} else {
			fmt.Println("test1 done")
			done <- true
			break
		}

	}

}

func main() {
	fmt.Println("test")

	items := make(chan int)
	done := make(chan bool)

	go test1(done, items, 1)
	go test1(done, items, 2)
	go test1(done, items, 3)

	i := 0
	for {
		items <- i
		i++
		if i > 10 {
			break
		}
	}

	close(items)
	var input string
	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println("Work done ... Press any key to exit")
	fmt.Scanln(&input)

}
