type unionFind struct {
	Parents []int
	Sizes   []int
}

func Constructor(length int) *unionFind {
	UF := &unionFind{
		Parents: make([]int, length),
		Sizes:   make([]int, length),
	}

	for i := 0; i < length; i++ {
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
	root1 := UF.find(ind1)
	root2 := UF.find(ind2)

	if root1 == root2 {
		return
	}

	if UF.Sizes[root1] > UF.Sizes[root2] {
		UF.Parents[root2] = root1
		UF.Sizes[root1] += UF.Sizes[root2]
	} else {
		UF.Parents[root1] = root2
		UF.Sizes[root2] += UF.Sizes[root1]
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	UF := Constructor(len(s))

	for _, pair := range pairs {
		UF.union(pair[0], pair[1])
	}

	// to connect all nodes of a root which are not connected yet
	for i := 0; i < len(s); i++ {
		UF.find(i)
	}

	// to add each union into a map to sort them
	unions := map[int][]byte{}
	for i := 0; i < len(s); i++ {
		root := UF.find(i)
		unions[root] = append(unions[root], s[i])
	}

	for key, val := range unions {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
		unions[key] = val
	}

	var res strings.Builder
	for i := 0; i < len(s); i++ {
		res.WriteByte(unions[UF.find(i)][0])
		unions[UF.find(i)] = unions[UF.find(i)][1:]
	}

	return res.String()
}