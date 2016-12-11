package main

import "fmt"

func main() {
	fmt.Println("Find factorial of a number using recursion")
	fmt.Println("Factorial of 7 = ", fact(7))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
