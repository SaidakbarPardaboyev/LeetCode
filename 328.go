/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	runner := head
	var bg *ListNode
	var runner2 *ListNode
	lamp := false

	for runner.Next != nil {
		var postNode *ListNode
		if runner.Next.Next != nil {
			postNode = runner.Next.Next
		} else {
			postNode = nil
		}
		if lamp {
			runner2.Next = runner.Next
			runner2 = runner2.Next
			runner2.Next = nil
		} else {
			lamp = true
			bg = runner.Next
			runner2 = runner.Next
			bg.Next = nil
		}
		if postNode != nil {
			runner.Next = postNode
			runner = runner.Next
		} else {
			break
		}
	}
	runner.Next = bg

	return head
}