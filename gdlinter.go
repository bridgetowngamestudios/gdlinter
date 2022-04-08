package main

import (
	"fmt"
	"os"

	"github.com/bridgetowngamestudios/gdlinter/pkg/linter"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Incorrect number of args passed.\nUsage: gdlinter <filename>\n")
		return
	}

	fmt.Printf("Linting file %s\n", args[1])

	err := linter.Lint(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
