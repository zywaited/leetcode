package one

import "sort"

func MaxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	start := 0
	prevIndex := start
	prevNum := nums[prevIndex]
	costK := 0
	ans := 1
	end := start + 1
	for end < len(nums) {
		for ; end < len(nums) && nums[end] == nums[start]; end++ {
		}
		costK += (nums[start] - prevNum) * (start - prevIndex)
		prevNum = nums[start]
		for costK > k {
			costK -= nums[start] - nums[prevIndex]
			prevIndex++
		}
		ans = max(ans, end-prevIndex)
		start = end
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
