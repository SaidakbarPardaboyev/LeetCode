/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	startNode := &ListNode{0, nil}
	endNode := &ListNode{Val: 0, Next: nil}
	st := startNode
	en := endNode

	for head != nil {
		prev := head.Next
		if head.Val >= x {
			endNode.Next = head
			endNode = endNode.Next
			endNode.Next = nil
		} else {
			startNode.Next = head
			startNode = startNode.Next
			startNode.Next = nil
		}
		head = prev
	}
	startNode.Next = en.Next
	return st.Next
}