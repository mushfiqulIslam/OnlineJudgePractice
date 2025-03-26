func countCompleteComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make(map[int]bool)
	completeCount := 0

	var dfs func(int, *[]int)
	dfs = func(node int, component *[]int) {
		if visited[node] {
			return
		}
		visited[node] = true

		*component = append(*component, node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor, component)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			var component []int
			dfs(i, &component)

			size := len(component)

			edgeCount := 0
			for _, v := range component {
				edgeCount = edgeCount + len(graph[v])

				if edgeCount == (size * (size - 1)) {
					completeCount = completeCount + 1
				}
			}
		}
	}
	return completeCount
}
