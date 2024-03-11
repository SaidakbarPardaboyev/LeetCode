package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	curNode := dummy
	NthPrev := dummy

	for i := 0; i < n; i++ {
		curNode = curNode.Next
	}

	for curNode.Next != nil {
		curNode = curNode.Next
		NthPrev = NthPrev.Next
	}
	NthPrev.Next = NthPrev.Next.Next

	return dummy.Next
}

func main() {
	var head *ListNode // Declare the head as a pointer to node

	// Add Node to the Linked List
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}

	AddToEnd(&head, node1)
	AddToEnd(&head, node2)
	AddToEnd(&head, node3)

	head = removeNthFromEnd(head, 3)

	// Print data of Linked list's Nodes
	PrintDataList(head)
}

// Print data of each node in linkedList
func PrintDataList(head *ListNode) {
	curNode := head
	for curNode != nil {
		fmt.Print(curNode.Val, " ")
		curNode = curNode.Next
	}
	fmt.Print("\n")
}

// Add to the end of the LinkedList
func AddToEnd(head **ListNode, newNode *ListNode) {
	if *head == nil {
		*head = newNode
		return
	}
	curNode := *head
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = newNode
}
