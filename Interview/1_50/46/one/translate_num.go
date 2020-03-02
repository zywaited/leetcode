package one

func TranslateNum(num int) int {
	dp := []int{1} // 默认1个
	pn := -1       // 前一个数字(0-9)
	cn := 0        // 当前数字
	i := 0
	for num > 0 {
		dp = append(dp, 0) // 初始化当前数据
		i++
		cn = num % 10
		num = num / 10
		dp[i] = dp[i-1]                           // 肯定和上次变化一样
		if cn != 0 && pn >= 0 && cn*10+pn <= 25 { // 可以结合[10-25]
			dp[i] += dp[i-2]
		}
		pn = cn
	}
	return dp[i]
}
