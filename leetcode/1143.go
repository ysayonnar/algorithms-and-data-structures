package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 0; i <= len(text1); i++ {
		dp[i][0] = 0
	}

	for i := 0; i <= len(text2); i++ {
		dp[0][i] = 0
	}

	lastFoundCol := 1
	for i := 1; i <= len(text1); i++ {
		for j := lastFoundCol; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}
