package leetcode

func countBits(n int) []int {
	bits := make([]int, n+1)
	bits[0] = 0
	d := 1
	for i := 1; i <= n; i++ {
		if i == d*2 {
			bits[i] = 1
			d *= 2
		} else {
			bits[i] = bits[i/2] + i%2
		}

	}

	return bits
}
