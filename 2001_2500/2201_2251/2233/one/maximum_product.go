package one

import "sort"

func MaximumProduct(nums []int, k int) int {
	mod := int(1e9 + 7)
	sort.Ints(nums)
	i := 0
	for ; k > 0; k-- {
		nums[i]++
		if i+1 == len(nums) || nums[i] <= nums[i+1] {
			i = 0
			continue
		}
		i++
	}
	ans := 1
	for _, num := range nums {
		ans = (ans * num) % mod
	}
	return ans
}
