package main

import "fmt"

func main() {
	fmt.Println("Using types")
	assertions()
	types()
}

func types() {
	typeSwitches(10)
	typeSwitches("hello")
	typeSwitches(nil)
	var test interface{} = "test"
	typeSwitches(test)
	typeSwitches(true)
}

func assertions() {
	var i interface{} = "hello"
	fmt.Printf("%v %T\n", i, i)

	s := i.(string)
	fmt.Printf("%v %T\n", s, s)

	// s holds the underlying value and ok holds a boolean value that
	// describes whether assertion is true or false
	s, ok := i.(string)
	fmt.Printf("%v %T %v\n", s, s, ok)

	// the below assertion is false as the string cannot be converted to float64
	// hence f will hold the zero value for type float64
	// ok will be false
	f, ok := i.(float64)
	fmt.Printf("%v %T %v\n", f, f, ok)

	// the below type assertion will trigger panic because we are not using the
	// assertion status returned by assertion
	//f = i.(float64)
	//fmt.Printf("%v %T\n", f, f)
}

func typeSwitches(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type is int")

	case string:
		fmt.Println("type is string")

	case nil:
		fmt.Println("type is nil")

	case bool:
		fmt.Println("type is boolean")

	default:
		fmt.Println("type is not known")

	}
}
