package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("panic")

	// panic aborts the program in case
	// we cannot handle the error that has occurred during any operation
	//panic("panic !!!")

	_, err := os.Create("/tmp/no-folder/error.txt")
	if err != nil {
		panic(err)
	}

}
