func restoreIpAddresses(s string) []string {
	res := []string{}

	var RestoreIPAdd func(ind int, curIP string, count int)
	RestoreIPAdd = func(ind int, curIP string, count int) {
		if ind >= len(s) {
			if count == 4 {
				res = append(res, curIP[:len(curIP)-1])
			}
			return
		}
		if count > 4 {
			return
		}

		for i := 0; i < 3 && ind+i < len(s); i++ {
			part := s[ind : ind+i+1]
			if isValidPart(part) {
				RestoreIPAdd(ind+i+1, curIP+part+".", count+1)
			}
		}
	}
	RestoreIPAdd(0, "", 0)

	return res
}

func isValidPart(part string) bool {
	if len(part) > 1 && part[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(part)
	return 0 <= num && num <= 255
}