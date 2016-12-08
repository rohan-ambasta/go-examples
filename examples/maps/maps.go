package main

import "fmt"

func main() {
	fmt.Println("Maps example")

	//to create a map use the built-in make
	m := make(map[string]int)

	//set key value pairs
	m["one"] = 1
	m["two"] = 2

	fmt.Println("map: ", m)

	//get value using the key
	v1 := m["one"]
	fmt.Println("value of one: ", v1)

	//print the length of the map
	fmt.Println("length of m: ", len(m))

	//Delete an entry from the map
	delete(m, "one")
	fmt.Println("map: ", m)

	//optional second return value while getting a value from the map
	//indicates whether the key is present or not. This can be used to disambiguate
	//between missing keys and keys with 0 or "" values. This can be also be used
	//in case we just need to check whether the key is present or not. Here
	//we don't need the value so we ignore it using the blank identifier
	_, isPresent := m["one"]
	fmt.Println("Is the value present: ", isPresent)

	//declare and initialize a map in the same line
	n := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	fmt.Println("map: ", n)
}
