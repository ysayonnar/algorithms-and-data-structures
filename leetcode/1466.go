package leetcode

func countPermutationsDfs(undirectedGraph map[int][]int, directedGraph map[int][]int, visited map[int]bool, current int) int {
	visited[current] = true
	ctr := 0

	for _, neighbour := range undirectedGraph[current] {
		if !visited[neighbour] {
			isFound := false
			for _, el := range directedGraph[neighbour] {
				if el == current {
					isFound = true
				}
			}
			if !isFound {
				ctr++
			}
			ctr += countPermutationsDfs(undirectedGraph, directedGraph, visited, neighbour)
		}
	}

	return ctr
}

func minReorder(n int, connections [][]int) int {
	visited := make(map[int]bool)

	undirectedGraph := make(map[int][]int)
	directedGrapg := make(map[int][]int)

	for _, edge := range connections {
		undirectedGraph[edge[0]] = append(undirectedGraph[edge[0]], edge[1])
		undirectedGraph[edge[1]] = append(undirectedGraph[edge[1]], edge[0])

		directedGrapg[edge[0]] = append(directedGrapg[edge[0]], edge[1])
	}

	return countPermutationsDfs(undirectedGraph, directedGrapg, visited, 0)
}
