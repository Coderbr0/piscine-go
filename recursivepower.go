package piscine

func RecursivePower(nb int, power int) int {
	var a int = 1

	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power >= 1 {
		a = nb * RecursivePower(nb, power-1)
	}
	return a
}
