package piscine

func IsAlpha(fraza string) bool {
	strbool := true
	for _, bukva := range fraza {
		if !(bukva >= 'a' && bukva <= 'z' || bukva >= 'A' && bukva <= 'Z' || bukva >= '0' && bukva <= '9') {
			strbool = false
		}
	}
	return strbool
}
