/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		length := len(queue)

		// Connect each levels nodes
		for i := 0; i < length-1; i++ {
			queue[i].Next = queue[i+1]
		}

		for i := 0; i < length; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}

	return root
}