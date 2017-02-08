package main

import (
	"fmt"
	"time"
)

func worker1(id int, jobs <-chan int, results chan<- int) {
	for {
		if job, more := <-jobs; more {
			fmt.Println("worker ", id, "started job ", job)
			time.Sleep(time.Second)
			fmt.Println("worker ", id, "finished job ", job)
			results <- job * 2
		} else {
			fmt.Println("All jobs complete")
			return
		}
	}

}

func main() {
	fmt.Println("Worker pool")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker1(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for k := 0; k < 5; k++ {
		fmt.Println(<-results)
	}

}
