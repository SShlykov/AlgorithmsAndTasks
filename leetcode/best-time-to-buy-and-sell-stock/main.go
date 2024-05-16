package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	buy := 0
	max_profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[buy] {
			buy = i
			continue
		}
		if prices[i]-prices[buy] > max_profit {
			max_profit = prices[i] - prices[buy]
		}
	}
	return max_profit
}
