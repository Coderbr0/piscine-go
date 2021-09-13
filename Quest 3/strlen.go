package piscine

func StrLen(s string) int {
	r := []rune(s)
	length := 0

	for index, letter := range r {
		letter += 0
		length = index

	}

	length += 1
	return length
}
