/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head.Next
	prev := head
	prev.Next = nil
	for cur != nil {
		tem := cur.Next
		cur.Next = prev
		prev = cur
		cur = tem
	}
	return prev
}