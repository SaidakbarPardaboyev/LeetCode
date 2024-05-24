/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	head := &Node{Val: 1, Neighbors: []*Node{}}
	visitedNodes := map[int]*Node{1: head}
	visited := map[int]bool{}
	var dfs func(curNode, CurResNode *Node)
	dfs = func(curNode, CurResNode *Node) {
		if visited[curNode.Val] {
			return
		}
		visited[curNode.Val] = true
		for i := 0; i < len(curNode.Neighbors); i++ {
			num := curNode.Neighbors[i].Val
			var nextNode *Node
			if _, ok := visitedNodes[num]; ok {
				nextNode = visitedNodes[num]
			} else {
				nextNode = &Node{Val: num, Neighbors: []*Node{}}
				visitedNodes[num] = nextNode
			}
			CurResNode.Neighbors = append(CurResNode.Neighbors, nextNode)
			dfs(curNode.Neighbors[i], nextNode)
		}
	}

	dfs(node, head)
	return head
}