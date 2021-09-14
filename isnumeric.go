package piscine

func IsNumeric(fraza string) bool {
	strbool := true
	for _, bukva := range fraza {
		if !(bukva >= '0' && bukva <= '9') {
			strbool = false
		}
	}
	return strbool
}
