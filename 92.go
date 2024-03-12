/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	runner := dummy
	var saveBegin *ListNode
	var saveFirst *ListNode
	var nextNode *ListNode
	var prevNode *ListNode
	count := 0
	lamp := false
	for runner != nil {
		nextNode = runner.Next
		if lamp {
			runner.Next = prevNode
		}
		if count == left-1 {
			saveBegin = runner
		} else if count == left {
			saveFirst = runner
			lamp = true
		} else if count == right {
			saveBegin.Next = runner
			saveFirst.Next = nextNode
			break
		}
		prevNode = runner
		if !lamp {
			runner = runner.Next
		} else {
			runner = nextNode
		}
		// fmt.Println(runner.Val)
		count++
	}
	return dummy.Next
}