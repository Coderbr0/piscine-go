package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 25 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := 1
	if nb > 0 {
		result = nb * RecursiveFactorial(nb-1)
		if result < 0 {
			return 0
		}
	}
	return result
}
