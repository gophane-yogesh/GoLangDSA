// Time Complexity O(n)
// Space Complexity O(1)

func maxProfit(prices []int) int {
	mProfit := 0
	prof := 0
	minValue := prices[0]

	for _, value := range prices[1:]{
		prof = value-minValue
		if prof > mProfit {
			mProfit=prof
		}
		if value<minValue{
			minValue=value
		}


	}

	return mProfit


}





