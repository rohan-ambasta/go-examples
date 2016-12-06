package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	argsWithoutProgram := os.Args[1:]

	command := exec.Command("go", "run", argsWithoutProgram[0]+".go")
	output, err := command.Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", output)
	}
}
