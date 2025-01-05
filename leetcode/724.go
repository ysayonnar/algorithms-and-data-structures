package leetcode

func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums))
	suffixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	suffixSum[len(nums)-1] = nums[len(nums)-1]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
		suffixSum[len(nums)-i-1] = suffixSum[len(nums)-i] + nums[len(nums)-i-1]
	}

	for i := 0; i < len(nums); i++ {
		if prefixSum[i] == suffixSum[i] {
			return i
		}
	}

	return -1
}
