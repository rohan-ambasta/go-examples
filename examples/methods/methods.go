package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

// methods are functions with a special argument 'receiver'
// below areaRect function has a receiver of type *rectangle
// *rectangle is a pointer receiver
func (r *rectangle) areaRect() float64 {
	return r.length * r.breadth
}

//the areaCircle method uses a value type receiver
//Go automatically handles conversion between values and pointers
//for method calls. You may want to use a pointer receiver type to
//avoid copying on method calls or to allow the method
//to mutate the receiving struct.
func (c circle) areaCircle() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

//You can declare a method on non-struct types, too.

//In this example we see a numeric type MyFloat with an Abs method.

//You can only declare a method with a receiver whose type is defined in
//the same package as the method. You cannot declare a method with a receiver
// whose type is defined in another package
//(which includes the built-in types such as int).

//MyFloat my float type
type MyFloat float64

//Abs method for MyFloat that returns absolute value
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//With a value receiver, the setRadius method operates on a copy
//of the original Circle value. (This is the same behavior as
//for any other function argument.) The setRadius method must have a pointer
//receiver to change the Radius.
func (c circle) setRadius(r float64) {
	c.radius = r
}

//Methods with pointer receivers can modify the value
//to which the receiver points. Since methods often need to
//modify their receiver, pointer receivers are more common than value receivers.
func (r *rectangle) setLength(l float64) {
	r.length = l
}

//BEST PRACTICE WHEN CHOOSING BETWEEN POINTER AND VALUE RECEIVERS

//There are two reasons to use a pointer receiver.

//The first is so that the method can modify the value that its receiver points to.

//The second is to avoid copying the value on each method call.
//This can be more efficient if the receiver is a large struct, for example.

//In general, all methods on a given type should have
//either value or pointer receivers, but not a mixture of both.

func main() {
	fmt.Println("Methods")
	rect := rectangle{10, 2}
	fmt.Println(rect)
	fmt.Println("area of rectangle = ", rect.areaRect())

	c := circle{2}
	fmt.Println(c)
	fmt.Println("area of circle = ", c.areaCircle())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	c.setRadius(3.0)
	fmt.Println(c)
	rect.setLength(5.0)
	fmt.Println(rect)

}
