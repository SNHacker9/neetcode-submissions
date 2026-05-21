func maxProfit(prices []int) int {
    maxProfit:=0 
	currProfit:=0

	for i:=0;i<len(prices)-1;i++{
		currProfit+= prices[i+1]-prices[i]

		if currProfit < 0 {
          currProfit = 0
		}

		if currProfit > maxProfit{
			maxProfit =currProfit
		}
	}

	return maxProfit
}

