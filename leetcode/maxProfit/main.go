package maxProfit

func maxProfit(prices []int) int {
	var sum int

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}

	return sum
}
