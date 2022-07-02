package one

func CountOrders(n int) int {
	mod := int(1e9 + 7)
	ans := 1
	dn := 0
	for i := n; i >= 1; i-- {
		ans = (ans * i) % mod
		ans = (ans * (dn + 1)) % mod
		dn += 2
	}
	return ans
}
