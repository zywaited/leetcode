package one

func MaximumSum(arr []int) int {
	al := len(arr)
	dp := make([][2]int, al)
	i := al - 1
	dp[i] = [2]int{arr[i], arr[i]}
	m := arr[i]
	i--
	for ; i >= 0; i-- {
		dp[i][0] = dp[i+1][0] + arr[i] // 一个都不去
		if arr[i] > dp[i][0] {
			dp[i][0] = arr[i] // 只保留自己
		}
		if dp[i][0] > m {
			m = dp[i][0]
		}
		dp[i][1] = dp[i+1][1] + arr[i] // 去掉别人
		if dp[i+1][0] > dp[i][1] {
			// 去掉自己更大
			dp[i][1] = dp[i+1][0]
		}
		if arr[i] > dp[i][1] {
			dp[i][1] = arr[i] // 只保留自己
		}
		if dp[i][1] > m {
			m = dp[i][1]
		}
	}
	return m
}
