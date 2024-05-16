type UnionFind struct {
	Parents []int
	Sizes   []int
}

func Constructor(length int) *UnionFind {
	UF := &UnionFind{
		Parents: make([]int, length),
		Sizes:   make([]int, length),
	}

	for i := 0; i < length; i++ {
		UF.Parents[i] = i
		UF.Sizes[i] = 1
	}
	return UF
}

func (UF *UnionFind) Find(ind int) int {
	if UF.Parents[ind] == ind {
		return ind
	}

	UF.Parents[ind] = UF.Find(UF.Parents[ind])
	return UF.Parents[ind]
}

func (UF *UnionFind) Union(ind1, ind2 int) {
	root1 := UF.Find(ind1)
	root2 := UF.Find(ind2)

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

func countServers(grid [][]int) int {
	UF := MakeGraph(grid)

	res := 0
	visited := map[int]bool{}
	for _, val := range UF.Parents {
		parent := UF.Find(val)
		if !visited[parent] && UF.Sizes[parent] > 1 {
			res += UF.Sizes[parent]
			visited[parent] = true
		}
	}

	return res
}

func MakeGraph(nums [][]int) *UnionFind {
	lengthOfRow := len(nums[0])
	lengthOfCol := len(nums)

	UF := Constructor(len(nums) * len(nums[0]))

	for i := 0; i < lengthOfCol; i++ {
		prev := -1
		for j := 0; j < lengthOfRow; j++ {
			if nums[i][j] == 1 {
				if prev != -1 {
					UF.Union(prev, i*lengthOfRow+j)
				} else {
					prev = i*lengthOfRow + j
				}
			}
		}
	}

	for j := 0; j < lengthOfRow; j++ {
		prev := -1
		for i := 0; i < lengthOfCol; i++ {
			if nums[i][j] == 1 {
				if prev != -1 {
					UF.Union(prev, i*lengthOfRow+j)
				} else {
					prev = i*lengthOfRow + j
				}
			}
		}
	}

	return UF
}