package leetcode

func tribonacci(n int) int {
	nums := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		nums = append(nums, nums[i-1]+nums[i-2]+nums[i-3])
	}
	return nums[n]
}
