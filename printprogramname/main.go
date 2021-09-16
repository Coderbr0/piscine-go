package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	a := argument[0][2:]

	for _, char := range a {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
