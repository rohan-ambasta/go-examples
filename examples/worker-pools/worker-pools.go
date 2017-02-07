package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, "started job ", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", job)
		results <- job * 2
	}

}

func main() {
	fmt.Println("Worker pools")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//for result := range results   {
	//	fmt.Println(result)
	//}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}

}
