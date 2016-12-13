package main

import "fmt"

//https://gobyexample.com/pointers
func main() {
	fmt.Println("Pointers example")

	i := 1
	fmt.Println("initial: ", i)

	zeroVal(i)
	fmt.Println("zeroval: ", i)

	zeroPtr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}
