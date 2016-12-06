package main

import (
	"fmt"
	"math"
)

//https://blog.golang.org/constants
//One way to think about untyped constants is that they live in a
//kind of ideal space of values, a space less restrictive than Go's full
//type system. But to do anything with them, we need to assign them to
//variables, and when that happens the variable (not the constant itself)
//needs a type, and the constant can tell the variable what type it should have.

//declare a constant using the const keyword
const s string = "constant"

func main() {
	fmt.Println(s)

	//A const statement can appear anywhere a var statement can.
	const n = 50000

	//constant of type string
	const typedHello string = "Hello, 世界"
	var s string
	s = typedHello
	fmt.Println(s)

	type MyString string
	var m MyString
	//cannot assign constant of type string to a variable of type MyString
	//m = typedHello // Type error

	//but an untyped constant can be assigned to variable of type MyString
	const unTypedHello = "Hello, 世界"
	m = unTypedHello
	fmt.Println(m)

	//The above example of typed and untyped constants also applies to other
	//types like int, int32, float, byte etc

	//Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	//A numeric constant has no type until it’s given one, such as by an explicit cast.
	fmt.Println(int64(d))

	//A number can be given a type by using it in a context
	//that requires one, such as a variable assignment or function call.
	//For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}
