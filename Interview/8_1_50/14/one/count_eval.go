package one

// DP
func CountEval(s string, result int) int {
	l := len(s)
	// 空也当作false
	if l == 0 {
		return 1 - result
	}
	// 实际数字长度
	n := (len(s) + 1) >> 1
	// i-j两种结果数量
	dp := make([][][2]int, n)
	// 索引
	fi, si, ti := 0, 0, 0
	for i := l - 1; i >= 0; i -= 2 {
		fi = i >> 1
		dp[fi] = make([][2]int, n)
		for j := i; j <= l-1; j += 2 {
			ti = j >> 1
			if i == j {
				if s[i] == '0' {
					dp[fi][ti][0] += 1
				} else {
					dp[fi][ti][1] += 1
				}
				continue
			}
			// 这里k不能等于j, 因为i-j加括号无意义，本来就是整体
			for k := i; k < j; k += 2 {
				si = k >> 1
				switch s[k+1] {
				case '&':
					dp[fi][ti][0] += dp[fi][si][0] * (dp[si+1][ti][0] + dp[si+1][ti][1])
					dp[fi][ti][0] += dp[fi][si][1] * dp[si+1][ti][0]
					dp[fi][ti][1] += dp[fi][si][1] * dp[si+1][ti][1]
				case '|':
					dp[fi][ti][0] += dp[fi][si][0] * dp[si+1][ti][0]
					dp[fi][ti][1] += dp[fi][si][1] * (dp[si+1][ti][0] + dp[si+1][ti][1])
					dp[fi][ti][1] += dp[fi][si][0] * dp[si+1][ti][1]
				case '^':
					dp[fi][ti][0] += dp[fi][si][0]*dp[si+1][ti][0] + dp[fi][si][1]*dp[si+1][ti][1]
					dp[fi][ti][1] += dp[fi][si][0]*dp[si+1][ti][1] + dp[fi][si][1]*dp[si+1][ti][0]
				}
			}
		}
	}
	return dp[0][n-1][result]
}
