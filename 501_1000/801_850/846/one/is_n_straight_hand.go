package one

import "sort"

func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	sort.Ints(hand)
	nums := make(map[int]int)
	for _, num := range hand {
		nums[num]++
	}
	for _, num := range hand {
		if nums[num] > 0 {
			for i := 1; i < groupSize; i++ {
				if nums[num+i] < nums[num] {
					return false
				}
				nums[num+i] -= nums[num]
			}
			nums[num] = 0
		}
	}
	return true
}
