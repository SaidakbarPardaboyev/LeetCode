func getSmallestString(s string, k int) string {
	i := 0
	answer := ""
	for i < len(s) && k > 0 {
		distanceBeforeLettertoA := int(s[i]) - int('a')
		distanceAfterLetterToA := 123 - int(s[i])
		if distanceBeforeLettertoA < distanceAfterLetterToA && distanceBeforeLettertoA <= k {
			answer += "a"
			k -= distanceBeforeLettertoA
		} else if distanceBeforeLettertoA > distanceAfterLetterToA && distanceAfterLetterToA <= k {
			answer += "a"
			k -= distanceAfterLetterToA
		} else if distanceBeforeLettertoA == distanceAfterLetterToA && distanceBeforeLettertoA <= k {
			answer += "a"
			k -= distanceAfterLetterToA
		} else {
			answer += string(int(s[i]) - k)
			k = 0
		}
		i++
	}
	answer += s[i:]
	return answer
}