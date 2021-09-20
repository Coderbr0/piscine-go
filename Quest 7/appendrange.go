package piscine

func AppendRange(min, max int) []int {
	var answer []int
	size := max - min
	if min >= max {
		return answer
	} else {
		if min < max {
			for i := 0; i < size; i++ {
				answer = append(answer, i+min)
			}
		}
	}
	return answer
}
