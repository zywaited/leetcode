package one

func NthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1
	indexes := make([]int, len(primes))
	for i := 1; i < n; i++ {
		mn := 0
		for j, index := range indexes {
			if mn == 0 || dp[index]*primes[j] < mn {
				mn = dp[index] * primes[j]
			}
		}
		dp[i] = mn
		for j, index := range indexes {
			if dp[index]*primes[j] == mn {
				indexes[j]++
			}
		}
	}
	return dp[n-1]
}
