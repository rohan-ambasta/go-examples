package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 100; i++ {
		a += 120 + i*40
	}

	fmt.Printf("%vKb\n", a/1024)
	fmt.Printf("size = %d\n", (256*4096)/1024)

}
