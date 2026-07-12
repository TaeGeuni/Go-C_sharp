package main

func countCompleteComponents(n int, edges [][]int) int {
	res := 0
	graph := make([][]int, n)

	for i := 0; i < len(edges); i++ {
		u := edges[i][0]
		v := edges[i][1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		queue := make([]int, 0)
		edgesNum := 0
		nodesNum := 0

		queue = append(queue, i)
		visited[i] = true

		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			nodesNum++

			edgesNum += len(graph[head])

			for j := 0; j < len(graph[head]); j++ {
				neighbor := graph[head][j]
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}

		if nodesNum*(nodesNum-1) == edgesNum {
			res++
		}
	}

	return res
}
