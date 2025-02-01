package leetcode

func maxSubArray(nums []int) int {
	curmax := nums[0]
	prev := nums[0]

	// я вахуе, первая dp которую полностью сам решил
	for i := 1; i < len(nums); i++ {
		prev = max(nums[i], nums[i]+prev)
		if prev > curmax {
			curmax = prev
		}
	}

	return curmax
}
