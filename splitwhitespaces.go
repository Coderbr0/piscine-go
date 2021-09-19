package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)  // break down string into runes
	var answer []string // create empty string array
	result := ""        // initialise result with empty string

	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' || runes[i] == '\n' || runes[i] == '\t' {
			if result != "" { // result updates every iteration
				answer = append(answer, result) // append previous result ie word
				result = ""                     // update result to be ""
			}
		} else {
			result = result + string(runes[i])
		}
		if i == len(runes)-1 {
			answer = append(answer, result) // append last result
		}
	}

	return answer
}
