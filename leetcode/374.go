package leetcode

func guess(num int) int { return 0 } //leetcode implements

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return 1
}
