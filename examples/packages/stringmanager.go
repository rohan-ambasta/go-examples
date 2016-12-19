package main

//"github.com/rohan-ambasta/go-practice/examples/packages/utils"
// read how to write go code https://golang.org/doc/code.html
import (
	"fmt"

	"github.com/rohan-ambasta/go-examples/examples/packages/utils"
)

func main() {
	fmt.Println(utils.Reverse(utils.ExportedVar))
}
