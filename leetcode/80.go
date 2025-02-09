package leetcode

func removeDuplicates(nums []int) int {
	writeIndex, ctr := 1, 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[readIndex-1] {
			ctr = 0
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		} else {
			ctr++
			if ctr <= 1 {
				nums[writeIndex] = nums[readIndex]
				writeIndex++
			}
		}
	}

	return writeIndex
}
