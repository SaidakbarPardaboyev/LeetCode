/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	mid := dummy
	end := dummy

	for true {
		if end.Next != nil {
			if end.Next.Next != nil {
				end = end.Next.Next
				mid = mid.Next
			} else {
				if mid.Next.Next != nil {
					mid.Next = mid.Next.Next
				} else {
					mid.Next = nil
				}
				break
			}
		} else {
			if mid.Next.Next != nil {
				mid.Next = mid.Next.Next
			} else {
				mid.Next = nil
			}
			break
		}
	}

	return dummy.Next
}