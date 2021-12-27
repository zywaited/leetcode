package one

func CombinationSum4(nums []int, target int) int {
	// 第一版提交按照正常的回溯最后一个实例会超时
	// 这一版动态规划
	dp := make(map[int]int, target+1)
	dp[0] = 1 // 空数组
	for tn := 1; tn <= target; tn++ {
		// 计算从1到target的组合数
		for _, num := range nums {
			dp[tn] += dp[tn-num]
		}
	}
	return dp[target]
}
