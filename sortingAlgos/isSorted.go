package sortingalgos

func IsSorted(arr []int) bool {
	n := len(arr)
	if n < 2 {
		return true
	}

	ascending := true
	descending := true

	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			ascending = false
		}
		if arr[i] > arr[i-1] {
			descending = false
		}
	}

	return ascending || descending
}
