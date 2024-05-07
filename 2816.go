/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	butun := 0

	var doubling func(node *ListNode)
	doubling = func(node *ListNode) {
		if node == nil {
			return
		}

		doubling(node.Next)
		tem := node.Val*2 + butun
		butun = (node.Val*2 + butun) / 10
		node.Val = tem % 10
	}
	doubling(head)

	if butun > 0 {
		newNode := &ListNode{
			Val:  butun,
			Next: head,
		}
		return newNode
	}
	return head
}