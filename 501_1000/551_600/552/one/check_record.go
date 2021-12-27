package one

func CheckRecord(n int) int {
	mod := int(1e9 + 7)
	curr := [2][3]int{}
	curr[0][0] = 1
	for i := 0; i < n; i++ {
		next := [2][3]int{}
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				if j < 1 {
					next[j+1][0] = (next[j+1][0] + curr[j][k]) % mod
				}
				if k < 2 {
					next[j][k+1] = (next[j][k+1] + curr[j][k]) % mod
				}
				next[j][0] = (next[j][0] + curr[j][k]) % mod
			}
		}
		curr, next = next, curr
	}
	ans := 0
	for _, jn := range curr {
		for _, kn := range jn {
			ans = (ans + kn) % mod
		}
	}
	return ans
}
