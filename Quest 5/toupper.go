package piscine

func ToUpper(fraza string) string {
	viwe_fraza := []rune(fraza)
	for index, bukva := range viwe_fraza {
		if bukva >= 'a' && bukva <= 'z' {
			viwe_fraza[index] = bukva - 32
		}
	}
	return string(viwe_fraza)
}
