package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	neg := n < 0
	if neg {
		n *= -1
	}
	returnSlice := []rune{}
	for {
		returnSlice = append([]rune{rune((n % 10) + '0')}, returnSlice...)
		if n/10 == 0 {
			break
		} else {
			n /= 10
		}
	}
	if neg {
		returnSlice = append([]rune{'-'}, returnSlice...)
	}
	for _, r := range returnSlice {
		z01.PrintRune(r)
	}
}
