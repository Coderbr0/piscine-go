package piscine

func ConcatParams(args []string) string {
	var s string
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			s = s + args[i]
		} else {
			s = s + args[i] + "\n"
		}
	}
	// return args[0] + "\n" +args[1] + "\n"+ args[2] + "\n"+ args[3]//
	return s
}
