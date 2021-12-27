package two

// 动态规划
// i,j代表从索引i到索引j的区间
// dp[i][j] = max(A[i] + dp[i+1][j]的后选值, A[j] + dp[i][j-1]的后选值)
// 所以当区间长度为1时，先手的值就是当前的数值，后手则为0，基于该逻辑计算
// 如：[5,3,4,5]
// dp[0][0] = [5, 0], dp[1][1] = [3, 0] ....
// dp[0][1] = max(5 + dp[1][1][1], 3 + dp[0][0][1])
// 刚好二维数组计算分布是斜的，就按该方向逐一计算
func StoneGame(piles []int) bool {
	// 二维数组dp[i][j]表示从i-j这段区间内先后手获取的最大分数
	type score struct {
		fir int // 先手
		sec int // 后手
	}
	num := len(piles)
	dp := make([][]score, num)
	for i := range dp {
		dp[i] = make([]score, num)
		// 单独的先手就是本身
		dp[i][i].fir = piles[i]
		dp[i][i].sec = 0
	}
	// 按照对称方向遍历(从1开始，0上面已经复制)
	for size := 1; size < num; size++ {
		for i := 0; i < num-size; i++ {
			j := size + i
			// i + (i+1 -- j的后手)
			l := piles[i] + dp[i+1][j].sec
			// j + (i -- j-1的后手)
			r := piles[j] + dp[i][j-1].sec
			if l >= r {
				dp[i][j].fir = l
				// 选择左边, (i+1 -- j)的先手变成当前区间的后手
				dp[i][j].sec = dp[i+1][j].fir
			} else {
				dp[i][j].fir = r
				dp[i][j].sec = dp[i][j-1].fir
			}
		}
	}
	// 取整个区间对比
	return dp[0][num-1].fir-dp[0][num-1].sec > 0
}
