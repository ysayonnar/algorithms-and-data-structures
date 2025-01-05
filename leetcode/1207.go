package leetcode

func uniqueOccurrences(arr []int) bool {
	set := make(map[int]int)
	occurences := make(map[int]bool)

	for _, num := range arr {
		if _, ok := set[num]; !ok {
			set[num] = 1
		} else {
			set[num]++
		}
	}

	for _, v := range set {
		occurences[v] = true
	}

	return len(occurences) == len(set)
}
