package one

const mod = 1e9 + 7

func ConcatenatedBinary(n int) int {
	ans := 0
	bn := uint(0)
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			// 二进制进1
			bn++
		}
		ans = ((ans << bn) + i) % mod
	}
	return ans
}
