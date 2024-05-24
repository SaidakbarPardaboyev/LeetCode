func MakeAdjacencyList(prerequisites [][]int) map[int][]int {
	adj := map[int][]int{}

	for _, val := range prerequisites {
		if _, ok := adj[val[0]]; !ok {
			adj[val[0]] = []int{}
		}
		adj[val[0]] = append(adj[val[0]], val[1])
	}

	return adj
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := MakeAdjacencyList(prerequisites)

	completedTasks := map[int]bool{}
	res := true
	var dfs func(node int, visited map[int]bool)
	dfs = func(node int, visited map[int]bool) {
		if completedTasks[node] {
			return
		}
		if visited[node] {
			res = false
			return
		}

		visited[node] = true
		for _, num := range adj[node] {
			dfs(num, visited)
		}
		completedTasks[node] = true
		delete(visited, node)
	}

	for i, _ := range adj {
		dfs(i, map[int]bool{})
		completedTasks[i] = true
	}

	return res
}