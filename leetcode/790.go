package leetcode

func numTilings(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % 1000000007
	}
	return dp[n]
}
