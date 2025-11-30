package one

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	numCount := make(map[int]int, len(nums))
	for _, num := range nums {
		numCount[num]++
	}
	leftIndex := -1
	rightIndex := 0
	var answer int
	for num := nums[0]; num < nums[len(nums)-1]; num++ {
		for ; num-nums[leftIndex+1] > k; leftIndex++ {
		}
		for ; rightIndex < len(nums) && nums[rightIndex]-num <= k; rightIndex++ {
		}
		answer = max(answer, max(rightIndex-leftIndex-2, numOperations+numCount[num]))
	}
	return answer
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
