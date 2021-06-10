package one

const mod = 1e9 + 7

func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][]int, minProfit+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := range profit {
		for p := minProfit; p >= 0; p-- {
			for g := n - group[i]; g >= 0; g-- {
				cp := p + profit[i]
				if cp > minProfit {
					cp = minProfit
				}
				dp[cp][g+group[i]] = (dp[cp][g+group[i]] + dp[p][g]) % mod
			}
		}
	}
	num := 0
	for _, p := range dp[minProfit] {
		num = (num + p) % mod
	}
	return num
}
