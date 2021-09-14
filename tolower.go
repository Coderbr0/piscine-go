package piscine

func ToLower(fraza string) string {
	viwe_fraza := []rune(fraza)
	for index, bukva := range viwe_fraza {
		if bukva >= 'A' && bukva <= 'Z' {
			viwe_fraza[index] = bukva + 32
		}
	}
	return string(viwe_fraza)
}
