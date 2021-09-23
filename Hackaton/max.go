package piscine

func Max(a []int) int {
	length := 0
	for l := range a {
		length = l + 1
	}
	i := 1
	for i < length {
		if a[i-1] > a[i] {
			temp := i
			a[i] = a[i-1]
			a[i-1] = a[temp]
			i = 1
		} else {
			i++
		}
	}

	return a[length-1]
}
