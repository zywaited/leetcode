package one

func minimumPossibleSum(n int, target int) int {
	mod := int(1e9 + 7)
	half := target / 2
	if n <= half {
		return (n * (n + 1) / 2) % mod
	}
	rn := n - half
	return (half*(half+1)/2 + target*rn + (rn-1)*rn/2) % mod
}
