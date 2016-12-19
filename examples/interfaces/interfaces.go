package main

import (
	"fmt"
	"math"
)

// Shape any geomtric shape
type Shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (rect *rectangle) area() float64 {
	return rect.length * rect.width
}

func (rect rectangle) perimeter() float64 {
	return 2*rect.length + 2*rect.width
}

func (c *circle) area() float64 {
	if c == nil {
		fmt.Printf("%v %T\n", c, c)
		return 0
	}
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s Shape) {
	//Under the covers, interface values can be thought of
	//as a tuple of a value and a concrete type: (value, type) as printed below
	fmt.Printf("%v %T\n", s, s)
	// in the below line area method will be called even if the pointer reciever
	// i.e the concrete shape value inside s is nil.
	// methods can be called with a nil receiver
	// But inside area method we will get nil pointer dereference error
	// while fetching the fields of that struct
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

//http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
//http://stackoverflow.com/questions/18125625/constructors-in-go
//https://golang.org/doc/effective_go.html#allocation_new
func main() {
	fmt.Println("Interfaces example")
	r := rectangle{2, 5}
	//will get the following error message while calling measure from line 51
	//cannot use r (type rectangle) as type shape in argument to measure:
	//rectangle does not implement shape (area method has pointer receiver)
	// the above error would not occur if we do not use pointer receivers for methods
	// but, not using pointer receivers is not a solution
	//measure(r)

	// the above error can be fixed by passing *rectangle to measure as below
	measure(&r)

	//instead of & operator we can also use new
	r1 := new(rectangle)
	r1.length = 2
	r1.width = 6
	measure(r1)

	//or you may use something like this
	c1 := &circle{5}
	measure(c1)

	// c1 is nil but the shape interface that holds c1 is non nil
	//c1 = nil
	//measure(c1)

	// in the below case the concrete type is nil
	// the interface methods will not get called because the type is not known
	measure(nil)
}
