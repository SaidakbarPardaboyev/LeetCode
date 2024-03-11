/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	curNode := dummy
	var last *ListNode
	count := -1
	for curNode != nil {
		if curNode.Next == nil {
			last = curNode
		}
		curNode = curNode.Next
		count++
	}

	if count < 2 {
		return dummy.Next
	}
	step := count - (k % count)
	if count == step {
		return dummy.Next
	}
	curNode = dummy
	header := dummy
	for step > 0 {
		curNode = curNode.Next
		step--
	}
	dummy = &ListNode{Val: 0, Next: curNode.Next}
	curNode.Next = nil
	last.Next = header.Next
	return dummy.Next
}