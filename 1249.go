func minRemoveToMakeValid(s string) string {
	stack := []rune{}
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else if s[i] == ')' {
			if len(stack) == 0 {
				s = s[:i] + s[i+1:]
				i--
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		i++
	}
	if len(stack) > 0 {
		i := len(s) - 1
		for len(stack) > 0 {
			if s[i] == '(' {
				s = s[:i] + s[i+1:]
				stack = stack[:len(stack)-1]
			}
			i--
		}
	}

	return s
}