package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Writing Files")

	// dump bytes into a file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	checkError(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/tmp/dat2")
	checkError(err)

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// write byte slices
	// write "some" into the file using a byte slice
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// writing string
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// flush writes to the disk using sync
	f.Sync()

	// write using bufio
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// flush the buffer into the file
	w.Flush()
}
