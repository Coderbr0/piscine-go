package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	aRune := "0123456789"

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if i == k && j == l {
						continue
					} else if k == i && l < j {
						continue
					} else if k < i {
						continue
					} else if (i == k && l > j) || (i != k || j != l) {

						z01.PrintRune(rune(aRune[i]))
						z01.PrintRune(rune(aRune[j]))

						z01.PrintRune(' ')

						z01.PrintRune(rune(aRune[k]))
						z01.PrintRune(rune(aRune[l]))

						if i == 9 && j == 8 && k == 9 && l == 9 {
							z01.PrintRune(rune(10))
							break
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
