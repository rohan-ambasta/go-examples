package main

import "fmt"

//https://blog.golang.org/slices

func main() {
	fmt.Println("Start working with slices")

	//initialize an empty array
	var emptyArray []int

	//initialize an array of strings
	stringValues := []string{"abc", "def", "xyz"}

	//Remember, since the slice header is always updated by a call to append,
	//you need to save the returned slice after the call. In fact, the compiler
	//won't let you call append without saving the result
	emptyArray = append(emptyArray, 10)
	emptyArray = append(emptyArray, 20)
	fmt.Printf("emptyArray = %v\n", emptyArray)

	stringValues = append(stringValues, "pqr", "lmn")
	var tempArray = append([]string{"uvw"}, "test")
	fmt.Printf("stringValues = %v\n", stringValues)
	fmt.Printf("tempArray = %v\n", tempArray)

	//append a slice
	//Not only can we append elements, we can append a whole second slice
	//by "exploding" the slice into arguments using the ... notation
	//at the call site:
	stringValues = append(stringValues, tempArray...)
	fmt.Println("Append stringValues and tempArray")
	fmt.Printf("stringValues = %v\n", stringValues)

	//copying a slice by appending to a nil slice
	//an empty slice can grow (assuming it has non-zero capacity),
	//but a nil slice has no array to put values in and can never
	//grow to hold even one element.

	//That said, a nil slice is functionally equivalent to a zero-length slice,
	//even though it points to nothing. It has length zero and can be appended to,
	//with allocation. As an example, look at the one-liner below that copies a
	//slice by appending to a nil slice.
	emptyArray1 := append([]int(nil), emptyArray...)
	fmt.Printf("emptyArray1 = %v\n", emptyArray1)

	slice1 := make([]int, 6, 10)
	slice1[0] = 1
	slice1[1] = 2
	fmt.Printf("slice1 = %v\n", slice1)

	//create slice2 from slice1
	slice2 := slice1[:2]
	fmt.Printf("slice2 = %v\n", slice2)
	slice2 = append(slice2, 3, 4)
	fmt.Println("Note that slice1 has also changed")
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice2 = %v\n", slice2)

	//to avoid the change in slice1
	//specify capacity while creating slice2 this will make sure
	//that a new array is created while appending values to slice2
	//as shown below using slice 3
	fmt.Println("create slice3 with capacity = 2")
	slice3 := slice1[:2:2]
	fmt.Printf("slice3 = %v\n", slice3)
	slice3 = append(slice3, 5, 6)
	fmt.Println("Note that slice1 does not change this time")
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("slice3 = %v\n", slice3)

	//Creating multi dimensional slices
	//Slices can be composed into multi-dimensional data structures.
	//The length of the inner slices can vary, unlike with multi-dimensional arrays.
	//https://gobyexample.com/slices
	fmt.Println("Creating multi dimensional slices ")
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//copying a slice using copy function
	fmt.Println("copying a slice using copy function")
	twoD1 := make([][]int, 3)
	copy(twoD1, twoD)
	fmt.Println("twoD1 = ", twoD1)

}
