package leetcode

func maxProfit(prices []int) int {
	profit, minimal := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if profit < prices[i]-minimal {
			profit = prices[i] - minimal
		} else if prices[i] < minimal {
			minimal = prices[i]
		}
	}

	return profit
}
