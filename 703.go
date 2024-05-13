type KthLargest struct {
	Heap []int
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := KthLargest{K: k}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	h.Push(0)

	for i := 0; i < k && i < len(nums); i++ {
		h.Push(nums[i])
	}

	return h
}

func (h *KthLargest) Delete(ind int) {
	for ind*2 < len(h.Heap) {
		if ind*2+1 < len(h.Heap) &&
			h.Heap[ind*2] > h.Heap[ind*2+1] &&
			h.Heap[ind*2+1] < h.Heap[ind] {
			h.Heap[ind], h.Heap[ind*2+1] = h.Heap[ind*2+1], h.Heap[ind]
			ind = ind*2 + 1
		} else if h.Heap[ind] > h.Heap[ind*2] {
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
		if h.Heap[ind] < h.Heap[ind/2] {
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
	if len(this.Heap) < this.K+1 {
		this.Push(val)
	} else {
		if val > this.Heap[1] {
			this.Heap[1] = val
			this.Delete(1)
		}
	}

	return this.Heap[1]

}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */