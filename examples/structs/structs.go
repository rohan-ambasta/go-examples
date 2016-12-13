package main

import "fmt"

//declare a struct with fields - name and age
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Structs example")

	//create a new struct
	fmt.Println("Create a new struct")
	fmt.Println(person{"Bob", 20})
	// compile time error
	//fmt.Println(person{20, "Bob"})

	fmt.Println("name the fields when initializing a struct")
	//you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})
	//no compile time error because order does not matter as  we are naming the fields
	fmt.Println(person{age: 10, name: "Jos"})

	//omitted fields can be zero valued
	fmt.Println("Omitted fields can be zero valued")
	fmt.Println(person{name: "ian"})
	fmt.Println(person{age: 10})

	fmt.Println("pointer to a struct using & prefix")
	fmt.Println(&person{name: "virat", age: 28})

	//access struct fields
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, s.age)

	//structs are mutable
	s.age = 21
	fmt.Println(s.age)

	//you can also use dot with struct pointers,
	// pointers are automatically dereferenced
	sp := &s
	sp.name = "John"
	fmt.Println(sp.name)
}
