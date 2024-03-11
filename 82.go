/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -101, Next: head}
	curNode := dummy

	var prevNode *ListNode
	// lamp := false

	for curNode.Next != nil {
		if curNode.Val == curNode.Next.Val {
			value := curNode.Val
			forcur := curNode
			for forcur.Val == value {
				forcur = forcur.Next
				if forcur == nil {
					prevNode.Next = nil
					break
				}
			}
			prevNode.Next = forcur
			curNode = prevNode
		} else {
			prevNode = curNode
			curNode = curNode.Next
		}
	}
	return dummy.Next
}