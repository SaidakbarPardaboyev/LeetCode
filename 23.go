/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	nums := addAllnumbersInList(lists)
	mergeSort(nums, 0, len(nums)-1)

	head := makeLinkedList(nums)

	return head
}

func addAllnumbersInList(lists []*ListNode) []int {
	nums := []int{}

	for _, head := range lists {
		curNode := head
		for curNode != nil {
			nums = append(nums, curNode.Val)
			curNode = curNode.Next
		}
	}

	return nums
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	left := make([]int, mid-start+1)
	copy(left, nums[start:mid+1])
	right := make([]int, end-mid)
	copy(right, nums[mid+1:end+1])

	l := 0
	r := 0
	i := start

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			nums[i] = left[l]
			l++
		} else {
			nums[i] = right[r]
			r++
		}
		i++
	}

	for l < len(left) {
		nums[i] = left[l]
		l++
		i++
	}
	for r < len(right) {
		nums[i] = right[r]
		r++
		i++
	}
}

func makeLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: 0}

	runner := head
	for _, val := range nums {
		runner.Next = &ListNode{Val: val}
		runner = runner.Next
	}

	return head.Next
}