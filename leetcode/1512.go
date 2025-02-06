package leetcode

func numIdenticalPairs(nums []int) int {
	ctr := 0

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == current {
				ctr++
			}
		}
	}

	return ctr
}
