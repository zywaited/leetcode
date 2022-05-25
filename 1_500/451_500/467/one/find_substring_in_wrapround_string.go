package one

func FindSubstringInWraproundString(p string) int {
	ans := 0
	bl := [26]int{}
	dp := 0
	for i := range p {
		j := int(p[i] - 'a')
		if i == 0 {
			dp = 1
			bl[j] = 1
			ans = 1
			continue
		}
		pl := bl[j]
		cl := 1
		pj := int(p[i-1] - 'a')
		if (pj+1)%26 == j {
			cl += dp
		}
		dp = cl
		if pl < cl {
			bl[j] = cl
			ans += cl - pl
		}
	}
	return ans
}
