/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	count := -1
	dummy := &ListNode{Val: 0, Next: head}

	runner := dummy
	for runner != nil {
		count++
		runner = runner.Next
	}

	if count < 2 {
		return
	}

	runner = dummy
	for i := 0; i < count/2+count%2; i++ {
		runner = runner.Next
	}
	beginNode := runner.Next
	runner.Next = nil

	right := reverseBetween(beginNode, 1, count/2)
	runner = right

	runner = dummy.Next
	runner2 := right
	var postNode *ListNode
	for runner2 != nil {
		postNode = runner2.Next
		runner2.Next = runner.Next
		runner.Next = runner2
		runner2 = postNode
		runner = runner.Next.Next
	}

	// return dummy
}

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