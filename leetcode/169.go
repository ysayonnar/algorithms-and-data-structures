package leetcode

func majorityElement(nums []int) int {
	numbers := make(map[int]int)
	n := len(nums)
	for _, num := range nums {
		numbers[num]++
		if numbers[num] > n/2 {
			return num
		}
	}

	return -1
}
