type Node struct {
	Val  string
	Next *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func ConstructorQueue() *Queue {
	return &Queue{}
}

type Stack struct {
	Head   *Node
	Length int
}

func ConstructorStack() *Stack {
	return &Stack{}
}

func (q *Queue) Enqueue(val string) {
	newNode := &Node{Val: val}
	if q.Tail == nil || q.Length == 0 {
		q.Tail = newNode
		q.Head = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = q.Tail.Next
	}
	q.Length++
}

func (q *Queue) Dequeue() (string, error) {
	if q.Length > 0 {
		res := q.Head.Val
		q.Head = q.Head.Next
		q.Length--
		return res, nil
	}
	return "", fmt.Errorf("Queue bo'sh")
}

func (s *Stack) Push(val string) {
	newNode := &Node{Val: val}
	s.Length++
	if s.Head == nil {
		s.Head = newNode
		return
	}
	newNode.Next = s.Head
	s.Head = newNode
}

func (s *Stack) Pop() (string, error) {
	if s.Length > 0 {
		res := s.Head.Val
		s.Head = s.Head.Next
		s.Length--
		return res, nil
	}
	return "", fmt.Errorf("Stack bo'sh")
}

func simplifyPath(path string) string {
	queue := ConstructorQueue()
	stack := ConstructorStack()

	for _, val := range strings.Split(path, "/") {
		if val == "." || val == "" {
			continue
		}
		queue.Enqueue(val)
	}

	for queue.Length > 0 {
		val, _ := queue.Dequeue()
		if val == ".." {
			stack.Pop()
		} else {
			stack.Push(val)
		}
	}

	res := ""
	for stack.Length > 0 {
		val, _ := stack.Pop()
		res = "/" + val + res
	}

	if len(res) == 0 {
		return "/"
	}
	return res
}