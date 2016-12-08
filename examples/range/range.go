package main

import "fmt"

func main() {
	fmt.Println("range example")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//iterate through numbers, ignore the index with blank identifier
	for _, number := range numbers {
		fmt.Println(number)
	}

	m := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}
	for key, value := range m {
		fmt.Println("key: ", key, " value: ", value)
	}

	//range on strings iterates over Unicode code points. The first value
	//is the starting byte index of the rune and the second the rune itself.
	for index, rune := range "golang" {
		fmt.Println("Starting bye index: ", index, " rune value: ", rune)
	}
}
