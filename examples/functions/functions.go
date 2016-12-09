package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Functions example")
	fmt.Println("2 + 3 = ", sum(2, 3))
	fmt.Println("10 - 2 = ", substract(10, 2))
	q, r := divide(10, 3)
	fmt.Println("10 / 3 = ", q, r)
	q1, r1 := divideFloat(13.44, 3.1)
	fmt.Println("13.44 / 3.1 = ", q1, r1)
	fmt.Println("Sum of integers in a slice = ", sumOfIntegers([]int{1, 2, 3, 4, 5}...))
}

//func that takes two int parameters and return the sum as int
func sum(a int, b int) int {
	return a + b
}

//When you have multiple consecutive arguments of the same type,
//you may omit the type name for the like-typed parameters up to
//the final parameter that declares the type.
func substract(a, b int) int {
	return a - b
}

//func with multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

//func with named return values
func divideFloat(a, b float64) (quotient, remainder float64) {
	quotient = math.Trunc(a / b)
	remainder = math.Mod(a, b)
	return
}

//variadic functions
//Variadic functions can be called with any number of trailing arguments.
//For example, fmt.Println is a common variadic function.
func sumOfIntegers(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
