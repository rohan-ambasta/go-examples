package main

import (
	"errors"
	"fmt"
)

type customError struct {
	errCode int
	errMsg  string
}

//It’s possible to use custom types as errors by implementing the Error() method on them.
func (e *customError) Error() string {
	return fmt.Sprintf("%s with error code %d", e.errMsg, e.errCode)
}

//By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg1 int) (int, error) {
	if arg1 == 10 {
		return 10 * 10, nil
	}
	//errors.New constructs a basic error value with the given error message.
	return -1, errors.New("The number is not 10")
}

func f2(arg int) (int, error) {
	if arg == 200 {
		return 200, nil
	}
	return 200, &customError{500, "some custom error"}
}

//https://blog.golang.org/error-handling-and-go
func main() {
	fmt.Println("errors")

	// A nil value in the error position indicates that there was no error.
	if _, e := f1(12); e != nil {
		fmt.Println("f1 failed: ", e)
	}

	if _, e := f2(12); e != nil {
		fmt.Println("f2 failed: ", e)
	}

	_, e := f2(12)
	//If you want to use the data
	// in a custom error, you’ll need to get the error as an instance
	//of the custom error type via type assertion.
	if custErr, ok := e.(*customError); ok {
		fmt.Println(custErr.errMsg)
	}
}
