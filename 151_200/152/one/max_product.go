package one

func MaxProduct(nums []int) int {
	ans := nums[0]
	// 保留上一次的最大&最小值
	maxN, minN := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxA, minA := maxN, minN
		maxN = max(nums[i], max(maxA*nums[i], minA*nums[i]))
		minN = min(nums[i], min(maxA*nums[i], minA*nums[i]))
		ans = max(ans, maxN)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
