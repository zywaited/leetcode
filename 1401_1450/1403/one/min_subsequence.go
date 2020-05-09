package one

import "sort"

func MinSubsequence(nums []int) []int {
	sort.Ints(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	ans := make([]int, 0, len(nums))
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		total -= nums[i]
		ans = append(ans, nums[i])
		if sum > total {
			break
		}
	}
	return ans
}
