package one

func KInversePairs(n int, k int) int {
	mod := 1000000007
	curr := make([]int, k+1)
	curr[0] = 1
	next := make([]int, k+1)
	for i := 1; i <= n; i++ {
		for j := k; j >= 0; j-- {
			next[j] = 0 // reset
			for l := max(i+j-k, 1); l <= i; l++ {
				next[i-l+j] = (next[i-l+j] + curr[j]) % mod
			}
		}
		curr, next = next, curr
	}
	return curr[k]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
