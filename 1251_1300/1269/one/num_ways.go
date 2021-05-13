package one

const mod = 1e9 + 7

func NumWays(steps int, arrLen int) int {
	pv := make([]int, steps+2)
	cv := make([]int, steps+2)
	pv[0] = 1
	for step := 1; step <= steps; step++ {
		for l := 0; l < arrLen; l++ {
			if l > step {
				// 根本过不去
				break
			}
			cv[l] = pv[l]
			if l > 0 {
				cv[l] += pv[l-1]
			}
			if l < arrLen-1 {
				cv[l] += pv[l+1]
			}
			cv[l] %= mod
		}
		pv, cv = cv, pv
	}
	return pv[0] % mod
}
