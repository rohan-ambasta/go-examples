package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("defer example")

	file := createFile("/tmp/temp.txt")
	// defer is like finally in java
	defer closeFile(file)
	writeFile(file)
}

func createFile(fileName string) *os.File {
	fmt.Println("Creating file")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("Writing to file")
	fmt.Fprintln(file, "Writing test data.")

}

func closeFile(file *os.File) {
	fmt.Println("closing file")
	file.Close()
}
