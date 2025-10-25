package one

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	numCount := make(map[int]int, len(nums))
	for _, num := range nums {
		numCount[num]++
	}
	var answer int
	for num := nums[len(nums)-1]; num >= nums[0]; num-- {
		leftIndex := sort.SearchInts(nums, num-k)
		rightIndex := sort.SearchInts(nums, num+k+1)
		currNumOperations := rightIndex - leftIndex - numCount[num]
		if numOperations < currNumOperations {
			currNumOperations = numOperations
		}
		count := currNumOperations + numCount[num]
		if answer < count {
			answer = count
		}
	}
	return answer
}
