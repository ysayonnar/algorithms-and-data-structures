package leetcode

func countRoomsDfs(rooms [][]int, visited map[int]bool, current int) int {
	visited[current] = true
	ctr := 1

	for _, neighbor := range rooms[current] {
		if !visited[neighbor] {
			ctr += countRoomsDfs(rooms, visited, neighbor)
		}
	}

	return ctr
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	ctr := countRoomsDfs(rooms, visited, 0)

	return len(rooms) == ctr
}
