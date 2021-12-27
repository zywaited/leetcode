package one

func NumTeams(rating []int) int {
	ans := 0
	// 中心扩散
	for j := 1; j < len(rating)-1; j++ {
		// 计算左右两边比当前大或小的数量
		ls, ll := 0, 0
		rs, rl := 0, 0
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				ls++
				continue
			}
			if rating[i] > rating[j] {
				ll++
			}
		}
		for k := j + 1; k < len(rating); k++ {
			if rating[j] < rating[k] {
				rl++
				continue
			}
			if rating[j] > rating[k] {
				rs++
			}
		}
		// 两边进行排列组合
		ans += ls*rl + ll*rs
	}
	return ans
}
