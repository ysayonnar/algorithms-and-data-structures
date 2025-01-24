package leetcode

func countDfs(isConnected [][]int, visited map[int]bool, current int) {
	visited[current] = true

	for neighbour, value := range isConnected[current] {
		if !visited[neighbour] && value == 1 {
			countDfs(isConnected, visited, neighbour)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	visited := make(map[int]bool)
	acc := 0

	//цикл + dfs для не связного графа
	//по сути задачи состоит в том, чтобы, найти количество связных графов в несвязном графе
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			countDfs(isConnected, visited, i)
			acc++
		}
	}

	return acc
}
