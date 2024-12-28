package main

import (
	sortingalgos "algos/sortingAlgos"
	"fmt"
)

func main() {
	arr := []int{1, 0, -1, -2, -3, 123, 2, 99}
	sortingalgos.BubbleSort(arr)
	fmt.Println(arr)
}
