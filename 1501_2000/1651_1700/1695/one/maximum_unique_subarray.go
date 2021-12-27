package one

func MaximumUniqueSubarray(nums []int) int {
	nm := make(map[int]int) // 数字对应的索引
	ans := 0
	sum := 0
	s := 0
	for i := 0; i < len(nums); i++ {
		if nm[nums[i]] > 0 {
			// 减去以前的数据
			for ; s < nm[nums[i]]; s++ {
				sum -= nums[s]
			}
		}
		sum += nums[i]
		nm[nums[i]] = i + 1
		ans = max(ans, sum)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
