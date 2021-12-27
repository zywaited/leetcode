package one

// 背包问题
func CanPartition(nums []int) bool {
	max := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	// 无法满足
	target := sum >> 1
	if sum&1 == 1 || target < max {
		return false
	}
	// 刚刚好
	if max == target {
		return true
	}
	dp := make([]bool, target+1)
	dp[0] = true // 0肯定没问题，空数组就算
	for _, num := range nums {
		// 选择当前数字
		// 从后开始循环是只用一次数字(从前开始会重复使用)
		for k := target; k >= num; k-- {
			dp[k] = dp[k] || dp[k-num]
		}
	}
	return dp[target]
}
