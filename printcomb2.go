package	piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	aRune := "0123456789"

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for 1 := 0; 1 < 10; 1++ {
					if i == k && j == 1   {
						continue
					} else if k == i && 1 < j {
						continue
					} else if k < i {
						continue
					} else if (i == k && 1 > j) || (i != k || j != 1) {

						z01.PrintRune(rune(aRune[i]))
						z01.PrintRune(rune(aRune[j]))

						z01.PrintRune(' ')

						z01.PrintRune(rune(aRune[k]))
						z01.PrintRune(rune(aRune[1]))

						if i == 9 && j == 8 && k == 9 && 1 == 9 {
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