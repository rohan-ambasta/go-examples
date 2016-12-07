package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Control structures")

	fmt.Println("If-Else")
	fmt.Println("Basic If-Else")
	//basic if else example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//You can have if statement without an else
	fmt.Println("You can have if statement without an else")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals;
	// Any variables declared in this statement
	// is available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, " is smaller than 10")
	} else {
		fmt.Println(num, " has multiple digits")
	}

	//Switch statements
	fmt.Println("")
	fmt.Println("Switch statments")

	//basic Switch, no need to explicitly write break statement after each case
	i := 5
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}

	//case can have multiple comma separated expressions
	// the default case is optional
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is a weekend")
	}

	//switch without expression is an alternative way to write if else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
