package leetcode

import (
	"sort"
)

// func removeThree(arr []int) []int {
// 	const maxOccurrences = 3
// 	elementCount := make(map[int]int)
// 	result := []int{}

// 	for _, num := range arr {
// 		if elementCount[num] < maxOccurrences {
// 			result = append(result, num)
// 			elementCount[num]++
// 		}
// 	}

// 	return result
// }

func isDuplicate(triplet []int, result [][]int) bool {
	for _, trip := range result {
		duplicate := false
		for i, t := range trip {
			if triplet[i] != t {
				duplicate = false
			} else {
				duplicate = true
			}
		}
		if duplicate {
			return true
		}
	}

	return false
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		current := nums[i]

		for current+nums[l]+nums[r] != 0 && l < r {
			sum := current + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else {
				l++
			}
		}

		if current+nums[l]+nums[r] == 0 && l != r {
			triplet := []int{current, nums[l], nums[r]}
			if !isDuplicate(triplet, result) {
				result = append(result, triplet)
			}
		}

	}

	return result
}
