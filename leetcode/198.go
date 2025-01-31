package leetcode

func rob(nums []int) int {
	prev1, prev2 := 0, 0
	for _, num := range nums {
		temp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = temp
	}

	return prev1
}
