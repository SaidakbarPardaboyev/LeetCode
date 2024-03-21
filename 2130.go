/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	if head.Next.Next == nil {
		return head.Val + head.Next.Val
	}

	mid := head
	end := head.Next
	for end.Next != nil {
		mid = mid.Next
		end = end.Next.Next
	}
	runner1 := head
	runner2 := reverseList(mid.Next)
	mid.Next = nil

	MaxSum := 0
	for runner1 != nil {
		if MaxSum < runner1.Val+runner2.Val {
			MaxSum = runner1.Val + runner2.Val
		}
		runner1 = runner1.Next
		runner2 = runner2.Next
	}

	return MaxSum
}

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