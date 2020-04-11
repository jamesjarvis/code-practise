func maxProfit(prices []int) int {
	sum := 0
	i := 0
	for i < (len(prices)-1) {
			
			for (i < (len(prices)-1)) && (prices[i+1] <= prices[i]) {
					i++
			}
			
			if i == (len(prices)-1) {
					return sum
			}
			
			buy := i
			i++
			
			for (i <= (len(prices)-1)) && (prices[i] >= prices[i-1]) {
					i++
			}
			sell := i-1
			
			sum += prices[sell]-prices[buy]
	}
	
	return sum
}