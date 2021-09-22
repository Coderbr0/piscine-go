package piscine

func Abort(a, b, c, d, e int) int {
	l := []int{a, b, c, d, e}
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l); j++ {
			for j := 0; j < len(l)-1; j++ {
				if l[j] > l[j+1] {
					l[j], l[j+1] = l[j+1], l[j]
				}
			}
		}
	}
	return l[2]
}
