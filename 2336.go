type SmallestInfiniteSet struct {
	Heap            []int
	ExistedNumber   map[int]bool
	LastNumberToPop int
}

func Constructor() SmallestInfiniteSet {
	h := SmallestInfiniteSet{
		Heap:            []int{0},
		ExistedNumber:   make(map[int]bool),
		LastNumberToPop: 1,
	}
	return h
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.Heap) == 1 || this.Heap[1] > this.LastNumberToPop {
		this.LastNumberToPop++
		return this.LastNumberToPop - 1
	} else if this.Heap[1] == this.LastNumberToPop {
		this.LastNumberToPop++
	}

	res := this.Heap[1]
	this.Heap[1] = this.Heap[len(this.Heap)-1]
	this.Heap = this.Heap[:len(this.Heap)-1]

	i := 1
	for i*2 < len(this.Heap) {
		if i*2+1 < len(this.Heap) &&
			this.Heap[i*2] > this.Heap[i*2+1] &&
			this.Heap[i] > this.Heap[i*2+1] {
			this.Heap[i], this.Heap[i*2+1] = this.Heap[i*2+1], this.Heap[i]
			i = i*2 + 1
		} else if this.Heap[i] > this.Heap[i*2] {
			this.Heap[i], this.Heap[i*2] = this.Heap[i*2], this.Heap[i]
			i = i * 2
		} else {
			break
		}
	}
	delete(this.ExistedNumber, res)
	return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.ExistedNumber[num] {
		return
	}
	this.Heap = append(this.Heap, num)

	i := len(this.Heap) - 1
	for i/2 >= 1 && this.Heap[i] < this.Heap[i/2] {
		this.Heap[i], this.Heap[i/2] = this.Heap[i/2], this.Heap[i]
		i /= 2
	}
	this.ExistedNumber[num] = true
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */