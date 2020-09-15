package one

func LongestOnes(A []int, K int) int {
	nl := 0
	tk := 0
	zq := make([]int, 0, K)
	pn := 0
	tn := 0
	for _, num := range A {
		if num == 1 {
			tn++
			continue
		}
		if tk < K {
			tk++
			tn++
			zq = append(zq, tn-pn)
			pn = tn
			continue
		}
		nl = max(nl, tn)
		if K == 0 {
			tn = 0
			continue
		}
		tn -= zq[0]
		pn -= zq[0]
		zq = zq[1:]
		tn++
		zq = append(zq, tn-pn)
		pn = tn
	}
	return max(nl, tn)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
