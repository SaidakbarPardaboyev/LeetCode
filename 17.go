func letterCombinations(digits string) []string {
	hash_table := map[byte][]byte{'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'}}
	res_comb := []string{}
	for len(digits) > 0 {
		cur := hash_table[digits[len(digits)-1]]
		digits = digits[:len(digits)-1]
		new_comb := []string{}
		if len(res_comb) > 0 {
			for i := 0; i < len(cur); i++ {
				for j := 0; j < len(res_comb); j++ {
					new_comb = append(new_comb, string(cur[i])+res_comb[j])
				}
			}
		} else {
			for _, val := range cur {
				new_comb = append(new_comb, string(val))
			}
		}
		res_comb = make([]string, len(new_comb))
		copy(res_comb, new_comb)
	}
	return res_comb
}