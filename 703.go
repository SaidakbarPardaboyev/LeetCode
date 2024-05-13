type KthLargest struct {
	Heap []int
	K    int
	MaxK int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{K: k}

	if len(nums) > 0 {
		nums = append(nums, nums[0])
		h.Heap = nums
	} else {
		h.Heap = []int{0}
	}

	for i := len(h.Heap) / 2; i > 0; i-- {
		h.Delete(i)
	}

	if len(nums) == 1 {
		res := -1
		copyOfHeap := append([]int{}, h.Heap...)
		for i := 0; i < h.K || i < len(h.Heap)-1; i++ {
			res = h.Pop()
		}
		h.MaxK = res
		h.Heap = append([]int{}, copyOfHeap...)
	} else {
		h.MaxK = math.MinInt64
	}

	return h
}

func (h *KthLargest) Delete(ind int) {
	for ind*2 < len(h.Heap) {
		if ind*2+1 < len(h.Heap) &&
			h.Heap[ind*2] < h.Heap[ind*2+1] &&
			h.Heap[ind*2+1] > h.Heap[ind] {
			h.Heap[ind], h.Heap[ind*2+1] = h.Heap[ind*2+1], h.Heap[ind]
			ind = ind*2 + 1
		} else if h.Heap[ind] < h.Heap[ind*2] {
			h.Heap[ind], h.Heap[ind*2] = h.Heap[ind*2], h.Heap[ind]
			ind *= 2
		} else {
			break
		}
	}
}

func (h *KthLargest) Push(val int) {
	if len(h.Heap) <= 1 {
		h.Heap = append(h.Heap, val)
		return
	}

	h.Heap = append(h.Heap, val)
	ind := len(h.Heap) - 1
	for ind/2 >= 1 {
		if h.Heap[ind] > h.Heap[ind/2] {
			h.Heap[ind], h.Heap[ind/2] = h.Heap[ind/2], h.Heap[ind]
			ind /= 2
		} else {
			break
		}
	}
}

func (h *KthLargest) Pop() int {
	res := h.Heap[1]
	if len(h.Heap) == 2 {
		h.Heap = []int{}
		return res
	}

	h.Heap[1] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]
	h.Delete(1)
	return res
}

func (this *KthLargest) Add(val int) int {
	this.Push(val)
	copyOfHeap := append([]int{}, this.Heap...)

	res := -1
	if val > this.MaxK {
		for i := 0; i < this.K; i++ {
			res = this.Pop()
		}
		this.Heap = append([]int{}, copyOfHeap...)
		this.MaxK = res
	} else {
		return this.MaxK
	}
	return res
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */