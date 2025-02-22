package leetcode

func r(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	r(nums, 0, n-1)
	r(nums, 0, k-1)
	r(nums, k, n-1)
}
