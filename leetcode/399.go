package leetcode

func calculateDfs(graph map[string]map[string]float64, current string, destination string, visited map[string]bool) float64 {
	visited[current] = true

	for neighbour, value := range graph[current] {
		if !visited[neighbour] {
			if neighbour == destination {
				return value
			} else {
				acc := calculateDfs(graph, neighbour, destination, visited)
				if acc != -1 {
					return value * acc
				}
			}
		}
	}
	return -1
}

// я ебал что это за пиздец а не задача ПОЧЕМУ ОНА МЕДИУМ Я РЕШАЛ ЕЕ ДВА ЧАСА НАХУЙ КАКИЕ УРАВНЕНИЯ ВЫ ЕБАНУЛИСЬ НА ЛИТКОДЕ ТАМ
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{} // откуда: кудa: вес

	for i := 0; i < len(equations); i++ {
		if _, ok := graph[equations[i][0]]; !ok {
			graph[equations[i][0]] = map[string]float64{}
		}
		if _, ok := graph[equations[i][1]]; !ok {
			graph[equations[i][1]] = map[string]float64{}
		}

		graph[equations[i][0]][equations[i][1]] = values[i]
		graph[equations[i][1]][equations[i][0]] = 1 / values[i]
	}

	result := []float64{}
	for _, query := range queries {
		if _, ok := graph[query[0]]; !ok {
			result = append(result, -1)
			continue
		} else if _, ok := graph[query[1]]; !ok {
			result = append(result, -1)
			continue
		}
		if query[0] == query[1] {
			result = append(result, 1)
			continue
		}
		visited := make(map[string]bool)
		result = append(result, calculateDfs(graph, query[0], query[1], visited))
	}

	return result
}
