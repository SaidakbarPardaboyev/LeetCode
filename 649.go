func predictPartyVictory(senate string) string {
	voting := senate
	// i = 0
	for len(voting) > 0 {
		lamp := false
		i := 1
		for i < len(voting) {
			if voting[0] != voting[i] {
				voting = voting[:i] + voting[i+1:]
				lamp = true
				break
			}
			i++
		}
		if !lamp {
			if voting[0] == 'R' {
				return "Radiant"
			} else {
				return "Dire"
			}
		}
		voting = voting[1:] + string(voting[0])
	}
	return ""
}