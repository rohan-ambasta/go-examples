package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sorting for builtin types")

	input := []string{"d", "a", "p"}
	// sorting is in place
	sort.Strings(input)
	fmt.Println("Sorted input - ", input)

	numInput := []int{10, 0, 1, 6, 100}
	sort.Ints(numInput)
	fmt.Println("sorted numbers - ", numInput)

	//sort in descending order
	// take the int slice and convert it into reverse type using
	// sort.Reverse. The reverse type is the input to the sort.Sort function.
	// The Reverse type implements the less method for sorting in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(numInput)))
	fmt.Println("numbers in reverse order - ", numInput)
}
