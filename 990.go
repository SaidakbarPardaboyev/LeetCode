type unionFind struct {
	Parents []int
	Sizes   []int
}

func Constructor() *unionFind {
	UF := &unionFind{
		Parents: make([]int, 26),
		Sizes:   make([]int, 26),
	}

	for i := 0; i < 26; i++ {
		UF.Parents[i] = i
		UF.Sizes[i] = 1
	}

	return UF
}

func (UF *unionFind) find(ind int) int {
	if ind == UF.Parents[ind] {
		return ind
	}

	UF.Parents[ind] = UF.find(UF.Parents[ind])
	return UF.Parents[ind]
}

func (UF *unionFind) union(ind1, ind2 int) {
	ultimatePar1 := UF.find(ind1)
	ultimatePar2 := UF.find(ind2)

	if ultimatePar1 == ultimatePar2 {
		return
	}

	if UF.Sizes[ultimatePar1] > UF.Sizes[ultimatePar2] {
		UF.Parents[ultimatePar2] = ultimatePar1
		UF.Sizes[ultimatePar1] += UF.Sizes[ultimatePar2]
	} else {
		UF.Parents[ultimatePar1] = ultimatePar2
		UF.Sizes[ultimatePar2] += UF.Sizes[ultimatePar1]
	}
}

func equationsPossible(equations []string) bool {
	UF := Constructor()
	restricted := make(map[byte][]byte)
	for _, val := range equations {
		if val[1:3] == "==" {
			checker := UF.find(int(val[3] - 'a'))
			for _, val := range restricted[val[0]] {
				if UF.find(int(val-'a')) == checker {
					return false
				}
			}
			checker = UF.find(int(val[0] - 'a'))
			for _, val := range restricted[val[3]] {
				if UF.find(int(val-'a')) == checker {
					return false
				}
			}
			UF.union(int(val[0]-'a'), int(val[3]-'a'))
		} else {
			restricted[val[0]] = append(restricted[val[0]], val[3])
			restricted[val[3]] = append(restricted[val[3]], val[0])
			if UF.find(int(val[0]-'a')) == UF.find(int(val[3]-'a')) {
				return false
			}
		}
	}
	return true
}