package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aStringChangeable := []byte(s)
	for i := 0; i < len(aStringChangeable); i++ {
		z01.PrintRune(rune(aStringChangeable[i]))
	}
}
