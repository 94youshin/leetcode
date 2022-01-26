func maxProfit(prices []int) int {
	profix := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profix += (prices[i+1]-prices[i])
		}
	}
	return profix
}
