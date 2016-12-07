package main

import "fmt"

func main() {
	fmt.Println("For is the only looping construct in Golang")

	fmt.Println("The most basic type, with a single condition.")
	i := 0
	//using for like a while loop
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//classic for loop
	fmt.Println("classic for loop example")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//infinite loop
	for {
		fmt.Println("Infinite Loop")
		break
	}

	//iterate over array/slice using for loop
	fmt.Println("Iterate over array/slice using for loop")

	numbers := []int{100, 200, 300, 400, 500}

	//standard iteration over array
	for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}

	//iterate like a foreach loop
	fmt.Println("Iterate using range in for loop")
	for index, number := range numbers {
		fmt.Printf("Index=%d , Value=%v\n", index, number)
	}

	//iterate like a foreach loop and ignore the index
	fmt.Println("Iterate using range in for loop and ignore the index")
	for _, number := range numbers {
		fmt.Printf("Value=%v\n", number)
	}
}
