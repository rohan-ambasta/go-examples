package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("Received job ", job)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job ", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
