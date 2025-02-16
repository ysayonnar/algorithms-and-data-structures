package leetcode

func mySqrt(x int) int {
	l, r := 0, x
	mid := (l + r) / 2

	for l < r {
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
		} else {
			return mid
		}

		mid = (l + r) / 2
	}

	if mid*mid > x {
		return mid - 1
	} else {
		return mid
	}
}
