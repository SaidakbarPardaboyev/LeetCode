type MyCircularQueue struct {
	Val      int
	prev     *MyCircularQueue
	Next     *MyCircularQueue
	lenth    int
	capacity int
}

// type ListNode struct {

// }

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		lenth:    0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsEmpty() {
		this.Val = value
		this.lenth++
		return true
	}
	if this.IsFull() {
		return false
	} else {
		if this.lenth < 2 {
			this.Next = &MyCircularQueue{
				Val:  value,
				prev: this,
				Next: this,
			}
			this.prev = this.Next
		} else {
			this.prev.Next = &MyCircularQueue{
				Val:  value,
				prev: this.prev,
				Next: this,
			}
			curPrev := this.prev.Next
			this.prev = curPrev
		}
		this.lenth++
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.lenth == 1 {
		this.Val = 0
		this.prev = nil
		this.Next = nil
		this.lenth = 0
	} else {
		nextNode := this.Next
		nextNode.prev = this.prev
		this.prev.Next = nextNode
		this.Val = nextNode.Val
		this.Next = nextNode.Next
		this.lenth--
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.lenth == 1 {
		return this.Val
	}
	return this.prev.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.lenth == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.lenth == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */