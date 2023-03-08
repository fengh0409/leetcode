package leetcode

// MaxProfit 买卖股票的最佳时机
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min = prices[0]
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		// 每次遍历，比较股票的最低价格
		if prices[i] < min {
			min = prices[i]
		}
		// 当前价格与最低股票之差即为利润，比较哪个利润高
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}
	return maxProfit
}
