package one

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfiedNum := 0
	for index, customer := range customers {
		if grumpy[index] == 0 {
			satisfiedNum += customer
		}
	}
	changedNum := 0
	customerIndex := 0
	for ; customerIndex < minutes; customerIndex++ {
		if grumpy[customerIndex] == 1 {
			changedNum += customers[customerIndex]
		}
	}
	maxChangedNum := changedNum
	for ; customerIndex < len(customers); customerIndex++ {
		if grumpy[customerIndex-minutes] == 1 {
			changedNum -= customers[customerIndex-minutes]
		}
		if grumpy[customerIndex] == 1 {
			changedNum += customers[customerIndex]
		}
		if maxChangedNum < changedNum {
			maxChangedNum = changedNum
		}
	}
	return satisfiedNum + maxChangedNum
}
