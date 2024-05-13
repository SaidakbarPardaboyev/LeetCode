type heap struct {
	Heap []int
}

func Heapify(nums []int) heap {
	heap := heap{}
	if len(nums) == 0 {
		heap.Heap = []int{0}
		return heap
	}

	heap.Heap = append([]int{0}, nums...)

	for i := len(heap.Heap) / 2; i > 0; i-- {
		heap.Reorder(i)
	}
	return heap
}

func (h *heap) Reorder(ind int) {
	for ind*2 < len(h.Heap) {
		if ind*2+1 < len(h.Heap) &&
			h.Heap[ind*2] < h.Heap[ind*2+1] &&
			h.Heap[ind] < h.Heap[ind*2+1] {
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

func (h *heap) Push(val int) {
	h.Heap = append(h.Heap, val)

	i := len(h.Heap) - 1
	for i/2 >= 1 {
		if h.Heap[i/2] < h.Heap[i] {
			h.Heap[i], h.Heap[i/2] = h.Heap[i/2], h.Heap[i]
			i /= 2
		} else {
			break
		}
	}
}

func (h *heap) Pop() int {
	res := h.Heap[1]

	h.Heap[1] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]

	h.Reorder(1)
	return res
}

func lastStoneWeight(stones []int) int {
	heap := Heapify(stones)

	for len(heap.Heap) > 2 {
		num1 := heap.Pop()
		num2 := heap.Pop()
		fmt.Println(num1)
		fmt.Println(num2)
		if num1-num2 > 0 {
			heap.Push(num1 - num2)
		}
	}

	if len(heap.Heap) <= 1 {
		return 0
	}
	return heap.Heap[1]
}