package sortingalgos

import "math/rand"

func Generate(length int) []int {
	arr := make([]int, length)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
