package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}

	result := 1

	for i := nb; i > 0; i-- {

		result = result * i

		if result < 0 {
			return 0
		}
	}
	return result
}
