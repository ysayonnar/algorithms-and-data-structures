package main

import (
	sortingalgos "algos/sortingAlgos"
	"fmt"
	"time"
)

func main() {
	length := 1000000000
	arr := sortingalgos.Generate(length)

	now := time.Now()
	arr = sortingalgos.MergeSort(arr)
	fmt.Println("Time elapsed: ", time.Since(now).Milliseconds(), "ms. Array length: ", length)
	fmt.Println("Is sorted: ", sortingalgos.IsSorted(arr))
}
