package _500_2000

func makeGood(s string) string {
	stack := []rune{}
	for _, r := range s {
		if len(stack) > 0 && (stack[len(stack)-1] == r+'a'-'A' || stack[len(stack)-1] == r+'A'-'a') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
