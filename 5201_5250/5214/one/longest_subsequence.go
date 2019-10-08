package one

func LongestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	max := 0
	for _, num := range arr {
		dp[num] = dp[num-difference] + 1
		if dp[num] > max {
			max = dp[num]
		}
	}
	return max
	// 最开始看错题目了，题目有要求是相邻的
	//dp := make(map[int]int)
	//for _, num := range arr {
	//	dp[num]++
	//}
	//max := 0
	//for num, count := range dp {
	//	if difference == 0 {
	//		if count > max {
	//			max = count
	//		}
	//		continue
	//	}
	//	if dp[num-difference] > 0 {
	//		// 最大最小的是其他，避免重复计算
	//		continue
	//	}
	//	tmp := 1
	//	for dp[num+difference] > 0 {
	//		tmp++
	//		num += difference
	//	}
	//	if tmp > max {
	//		max = tmp
	//	}
	//}
	//return max
}
