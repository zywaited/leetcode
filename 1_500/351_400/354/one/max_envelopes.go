package one

import "sort"

func MaxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([]int, len(envelopes))
	ans := 0
	for j := 0; j < len(envelopes); j++ {
		dp[j] = 1
		for i := j - 1; i >= 0; i-- {
			if envelopes[j][0] > envelopes[i][0] && envelopes[j][1] > envelopes[i][1] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		ans = max(ans, dp[j])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
