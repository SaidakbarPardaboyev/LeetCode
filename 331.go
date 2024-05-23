func isValidSerialization(preorder string) bool {
	items := strings.Split(preorder, ",")
	stack := []int{}
	if items[0] == "#" && len(items) > 1 {
		return false
	}
	for _, val := range items {
		if len(stack) != 0 {
			stack[len(stack)-1]++
			if stack[len(stack)-1] == 2 {
				stack = stack[:len(stack)-1]
			}
		} else if val != items[0] {
			return false
		}
		if val != "#" {
			stack = append(stack, 0)
		}
	}
	return len(stack) == 0
}