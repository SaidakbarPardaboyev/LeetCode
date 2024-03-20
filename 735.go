type stack []int

func (s *stack) Add(num int) {
	if len(*s) == 0 {
		*s = append(*s, num)
		return
	}
	i := len(*s) - 1
	for i >= 0 {
		if (*s)[len(*s)-1] > 0 && num < 0 {
			num2 := -num
			if (*s)[i] == num2 {
				*s = (*s)[:i]
				break
			} else if (*s)[i] < num2 {
				*s = (*s)[:i]
				if len(*s) == 0 {
					*s = append(*s, num)
					break
				}
			} else {
				break
			}
		} else {
			*s = append(*s, num)
			break
		}
		i--
	}
}

func asteroidCollision(asteroids []int) []int {
	stack := stack{}
	for _, astroid := range asteroids {
		stack.Add(astroid)
	}
	return stack
}