/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	if head == nil {
		return head
	}
	cur := head
	for cur != nil {
		if visited[cur] {
			return cur
		}
		visited[cur] = true
		cur = cur.Next
	}
	return nil
}