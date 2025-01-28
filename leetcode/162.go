package leetcode

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2

		if mid == 0 {
			if nums[mid] > nums[mid+1] {
				return mid
			} else {
				l = mid + 1
				continue
			}
		} else if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			} else {
				r = mid - 1
				continue
			}
		}

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid-1] > nums[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
