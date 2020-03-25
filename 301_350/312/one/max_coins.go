package one

func MaxCoins(nums []int) int {
	// 区间计算最大值
	return block(nums, 0, len(nums)-1, make([][]int, len(nums)))
}

// 戳破s-e区间的气球
func block(nums []int, s, e int, cache [][]int) int {
	if s > e {
		return 0
	}
	if cache[s] == nil {
		cache[s] = make([]int, len(nums))
	}
	if cache[s][e] > 0 {
		return cache[s][e]
	}
	coins := 0
	for cs := s; cs <= e; cs++ {
		coins = nums[cs]
		if s-1 >= 0 {
			coins *= nums[s-1]
		}
		if e+1 < len(nums) {
			coins *= nums[e+1]
		}
		coins += block(nums, s, cs-1, cache) + block(nums, cs+1, e, cache)
		cache[s][e] = max(cache[s][e], coins)
	}
	return cache[s][e]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
