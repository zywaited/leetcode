package one

func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if boardingCost*4 <= runningCost {
		return -1
	}
	times := -1
	maxProfit := 0
	customerNum := 0
	customerIndex := 0
	currentTimes := 0
	currentProfit := 0
	for customerIndex < len(customers) || customerNum > 0 {
		if customerIndex < len(customers) && customerIndex <= currentTimes {
			customerNum += customers[customerIndex]
			customerIndex++
			continue
		}
		if customerNum == 0 {
			currentProfit -= runningCost
			currentTimes++
			continue
		}
		canTimes := customerNum / 4
		if canTimes > 0 {
			currentProfit += (boardingCost*4 - runningCost) * canTimes
			currentTimes += canTimes
			customerNum %= 4
		} else {
			currentProfit += boardingCost*customerNum - runningCost
			currentTimes++
			customerNum = 0
		}
		if maxProfit < currentProfit {
			times = currentTimes
			maxProfit = currentProfit
		}
	}
	return times
}
