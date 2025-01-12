package sortingalgos

func merge(left, right []int) []int {
	res := []int{}

	for len(left) != 0 && len(right) != 0 {
		if left[0] > right[0] {
			res = append(res, right[0])
			right = right[1:]
		} else {
			res = append(res, left[0])
			left = left[1:]
		}
	}

	if len(left) == 0 && len(right) != 0 {
		res = append(res, right...)
	} else if len(left) != 0 && len(right) == 0 {
		res = append(res, left...)
	}
	return res
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}
