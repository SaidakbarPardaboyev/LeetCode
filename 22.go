func generateParenthesis(n int) []string {
	res := []string{}

	var generateParentheses func(open, close int, parenthes string)
	generateParentheses = func(open, close int, parenthes string) {
		if open == n && close == n {
			res = append(res, parenthes)
			return
		}

		if open < n {
			generateParentheses(open+1, close, parenthes+"(")
		}
		if open > close {
			generateParentheses(open, close+1, parenthes+")")
		}
	}
	generateParentheses(0, 0, "")

	return res
}