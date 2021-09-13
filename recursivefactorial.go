package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := 1

	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return result
}
