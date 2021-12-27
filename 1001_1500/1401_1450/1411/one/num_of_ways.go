package one

const mod = 1e9 + 7

func NumOfWays(n int) int {
	// 一行的颜色要么使用两色，要么使用三色

	// 第一行使用两色和三色各有6种情况
	// un 代表上一行
	un := [2]int{6, 6}
	// cn 代表当前行
	cn := [2]int{}
	for row := 1; row < n; row++ {
		// 上一行: ABA | ABC
		// 下一行两色: B[AC]B、CAC 一共3种情况, B[AC]B 一共2种情况
		// 下一行三色: BAC、CAB 一共2种情况, BCA、CAB一共2种情况
		cn[0] = int((un[0]*3)%mod+(un[1]*2)%mod) % mod
		cn[1] = int((un[0]*2)%mod+(un[1]*2)%mod) % mod
		un, cn = cn, un
	}
	return (un[0] + un[1]) % mod
}
