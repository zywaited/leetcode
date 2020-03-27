package one

func HasGroupsSizeX(deck []int) bool {
	dn := make(map[int]int, len(deck))
	for _, d := range deck {
		dn[d]++
	}
	// 最大公约数
	g := 0
	for _, num := range dn {
		if g == 1 {
			// 不用操作了
			break
		}
		if g == 0 {
			g = num
			continue
		}
		g = gcd(g, num)
	}
	return g >= 2
}

func gcd(f, s int) int {
	for s > 0 {
		f, s = s, f%s
	}
	return f
}
