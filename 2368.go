func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := MakeGraph(edges)
	visited := make(map[int]bool)
	for _, val := range restricted {
		visited[val] = true
	}
	queue := []int{0}
	NodesCount := 0

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			visited[cur] = true
			NodesCount++
			for _, neigh := range graph[cur] {
				if !visited[neigh] {
					queue = append(queue, neigh)
				}
			}
		}
	}
	return NodesCount
}

func MakeGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, neighbours := range edges {
		graph[neighbours[0]] = append(graph[neighbours[0]], neighbours[1])
		graph[neighbours[1]] = append(graph[neighbours[1]], neighbours[0])
	}
	return graph
}