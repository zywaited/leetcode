package one

import "sort"

func minOperations(nums []int) int {
	sort.Ints(nums)
	var uniqueNums []int
	for index := range nums {
		if index == 0 || nums[index] != nums[index-1] {
			uniqueNums = append(uniqueNums, nums[index])
		}
	}
	answer := len(nums)
	rightIndex := 0
	for leftIndex := range uniqueNums {
		for rightIndex < len(uniqueNums) && uniqueNums[rightIndex]-uniqueNums[leftIndex]+1 <= len(nums) {
			rightIndex++
		}
		answer = min(answer, leftIndex+(len(uniqueNums)-rightIndex)+(len(nums)-len(uniqueNums)))
	}
	return answer
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
