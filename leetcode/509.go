package leetcode

func fib(n int) int {
	if n < 2 {
		return n
	}

	prev, cur := 0, 1

	for i := 1; i < n; i++ {
		prev, cur = cur, prev+cur
	}

	return cur
}
