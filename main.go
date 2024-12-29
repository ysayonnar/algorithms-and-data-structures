package main

import (
	sortingalgos "algos/sortingAlgos"
	"fmt"
)

func main() {
	arr := []int{1, 0, -1, -2, -3, 123, 2, 99}
	fmt.Println("Before: ", arr)
	sortingalgos.SelectionSort(arr)
	fmt.Println("After: ", arr)

	fmt.Println("Is sorted: ", sortingalgos.IsSorted(arr))
}
