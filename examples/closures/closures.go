package main

import "fmt"

//useful links for knowing more about closures
// 1. http://www.calhoun.io/what-is-a-closure/
// 2. http://www.calhoun.io/5-useful-ways-to-use-closures-in-go/
func main() {
	fmt.Println("Start with closures")

	nextInt := intSeq(0)
	fmt.Println("nextInt to 0 is ", nextInt())
	fmt.Println("nextInt to 1 is ", nextInt())
	fmt.Println("nextInt to 2 is ", nextInt())

	nextInt1 := intSeq(10)
	fmt.Println("nextInt to 10 is ", nextInt1())

	//print fibonacci series
	fmt.Println("print fibonacci series")
	fiboGenerator()
}

func fiboGenerator() {
	gen := makeFibSeq()

	for index := 0; index < 10; index++ {
		fmt.Println(gen(), " ")
	}
}

//return a function
func intSeq(a int) func() int {
	//anonymous function that returns int
	// this function captures its own value of a and increments it
	return func() int {
		a++
		return a
	}
}

//make fibonacci sequence
func makeFibSeq() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = f1+f2, f2
		return f1
	}
}
