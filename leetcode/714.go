package leetcode

func maxProfit2(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}

	hold := -prices[0]
	cash := 0

	for _, price := range prices {
		newHold := max(hold, cash-price)
		newCash := max(cash, hold+price-fee)
		hold, cash = newHold, newCash
	}

	return cash
}
