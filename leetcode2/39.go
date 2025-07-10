package leetcode2

func tryToFind(candidates []int, target int) ([]int, bool) {
	if target == 0 {
		return []int{}, true
	} else if target < 0 {
		return []int{}, false
	}

	for _, candidate := range candidates {
		combinaiton, isFound := tryToFind(candidates, target-candidate)
		if isFound {
			combinaiton = append(combinaiton, candidate)
			return combinaiton, true
		}
	}

	return []int{}, false
}

func combinationSum(candidates []int, target int) [][]int {
	combinations := [][]int{}
	for _, candidate := range candidates {
		combination, isFound := tryToFind(candidates, target-candidate)
		if isFound {
			combination = append(combination, candidate)
			combinations = append(combinations, combination)
		}
	}

	return combinations
}
