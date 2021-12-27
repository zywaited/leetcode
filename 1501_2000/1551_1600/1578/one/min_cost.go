package one

func MinCost(s string, cost []int) int {
	cs := 0
	index := 0
	num := 0
	sum := 0
	maxIndex := 0
	for index < len(s) {
		maxIndex = index
		num = 0
		sum = 0
		for ; index < len(s) && s[index] == s[maxIndex]; index++ {
			num++
			sum += cost[index]
			if num > 1 && cost[index] > cost[maxIndex] {
				maxIndex = index
			}
		}
		if num > 1 {
			cs += sum - cost[maxIndex]
		}
	}
	return cs
}
