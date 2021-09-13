package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, num := range s {
		if num >= 48 && num <= 57 {
			conv := int(num) - 48
			result = (result * 10) + conv
		} else {
			return 0
		}
	}
	return result
}
