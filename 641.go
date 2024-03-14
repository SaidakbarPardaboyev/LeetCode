package main

import "fmt"

func main() {
	// Your MyCircularDeque object will be instantiated and called as such:
	obj := Constructor(3)
	param_1 := obj.InsertLast(1) // null
	fmt.Println(param_1)
	param_2 := obj.InsertLast(2) // true
	fmt.Println(param_2)
	param_3 := obj.InsertFront(3) // true
	fmt.Println(param_3)
	param_4 := obj.InsertFront(4) // false
	fmt.Println(param_4)
	param_5 := obj.GetRear() // 2
	fmt.Println(param_5)
	param_6 := obj.IsFull() // true
	fmt.Println(param_6)
	param_7 := obj.DeleteLast() // true
	fmt.Println(param_7)
	param_8 := obj.InsertFront(4) // true
	fmt.Println(param_8)
	param_9 := obj.GetFront() // 4
	fmt.Println(param_9)
	// param_7 := obj.IsEmpty()
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func MakeNode(Val int, PreNode, PostNode *Node) *Node {
	return &Node{
		Val:  Val,
		Prev: PreNode,
		Next: PostNode,
	}
}

type MyCircularDeque struct {
	head     *Node
	Last     *Node
	Length   int
	Capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Length:   0,
		Capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.Length == 0 {
		this.head = MakeNode(value, this.head, this.head)
	} else if this.Length == 1 {
		New := MakeNode(value, this.head, this.head)
		this.head.Prev = New
		this.head.Next = New
		this.Last = this.head
		this.head = New
	} else {
		New := MakeNode(value, this.Last, this.head)
		this.Last.Next = New
		this.head.Prev = New
		this.head = New
	}
	this.Length++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.Length == 0 {
		this.head = MakeNode(value, this.head, this.head)
	} else if this.Length == 1 {
		New := MakeNode(value, this.head, this.head)
		this.Last = New
		this.head.Next = New
		this.head.Prev = New
	} else {
		New := MakeNode(value, this.Last, this.head)
		this.Last.Next = New
		this.head.Prev = New
		this.Last = New
	}
	this.Length++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Length == 1 {
		this.head = nil
	} else {
		this.head.Next.Prev = this.Last
		this.Last.Next = this.head.Next
		this.head = this.head.Next
	}
	this.Length--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Length == 1 {
		this.head = nil
	} else {
		this.head.Prev = this.Last.Prev
		this.Last.Prev.Next = this.head
		this.Last = this.Last.Prev
	}
	this.Length--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.Length == 0 {
		return -1
	}
	return this.head.Val
}

func (this *MyCircularDeque) GetRear() int {
	if this.Length == 0 {
		return -1
	}
	if this.Length == 1 {
		return this.head.Val
	}
	return this.Last.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Length == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Length == this.Capacity
}
