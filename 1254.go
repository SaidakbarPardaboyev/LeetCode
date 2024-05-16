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

func closedIsland(grid [][]int) int {
	lengthOfRow := len(grid[0])
	lengthOfCol := len(grid)

	UF := MakeGraph(grid)
	restricted := MakeRestricted(UF, grid)

	res := map[int]bool{}

	for i := 1; i < lengthOfCol-1; i++ {
		for j := 1; j < lengthOfRow-1; j++ {
			if grid[i][j] == 0 {
				index := UF.Find(i*lengthOfRow + j)
				if !restricted[UF.Parents[index]] {
					res[UF.Parents[index]] = true
				}
			}
		}
	}
	return len(res)
}

func MakeGraph(nums [][]int) *UnionFind {
	lengthOfRow := len(nums[0])
	lengthOfCol := len(nums)
	UF := Constructor(lengthOfRow * lengthOfCol)

	for i := 0; i < lengthOfCol; i++ {
		for j := 0; j < lengthOfRow; j++ {
			if nums[i][j] == 1 {
				continue
			}
			neighbours := FindNeighbours(i, j, nums)
			for _, xy := range neighbours {
				UF.Union(i*lengthOfRow+j, xy[0]*lengthOfRow+xy[1])
			}
		}
	}
	return UF
}

func FindNeighbours(i, j int, nums [][]int) [][]int {
	lengthOfRow := len(nums[0])
	lengthOfCol := len(nums)
	res := [][]int{}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for _, direction := range directions {
		row := i + direction[0]
		col := j + direction[1]
		if row >= 0 && row < lengthOfCol &&
			col >= 0 && col < lengthOfRow &&
			nums[i][j] == nums[row][col] {
			res = append(res, []int{row, col})
		}
	}
	return res
}

func MakeRestricted(UF *UnionFind, nums [][]int) map[int]bool {
	lengthOfRow := len(nums[0])
	lengthOfCol := len(nums)

	res := map[int]bool{}

	// for First and Last Row
	for i := 0; i < lengthOfRow; i++ {
		if nums[0][i] == 0 {
			res[UF.Find(i)] = true
		}
		if nums[lengthOfCol-1][i] == 0 {
			res[UF.Find((lengthOfCol-1)*lengthOfRow+i)] = true
		}
	}
	// for First and Last Column
	for i := 0; i < lengthOfCol; i++ {
		if nums[i][0] == 0 {
			res[UF.Find(i*lengthOfRow)] = true
		}
		if nums[i][lengthOfRow-1] == 0 {
			res[UF.Find(i*lengthOfRow+lengthOfRow-1)] = true
		}
	}
	return res
}