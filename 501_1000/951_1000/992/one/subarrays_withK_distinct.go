package one

func SubarraysWithKDistinct(A []int, K int) int {
	pn := 0
	pi := 0
	nk := 0
	an := make([]int, len(A)+1)
	ans := 0
	for _, a := range A {
		an[a]++
		if an[a] == 1 {
			nk++
		}
		if nk < K {
			continue
		}
		if nk > K {
			for nk > K {
				an[A[pi]]--
				if an[A[pi]] == 0 {
					nk--
				}
				pi++
			}
			pn = 0
		}
		for nk == K && an[A[pi]] > 1 {
			an[A[pi]]--
			pn++
			pi++
		}
		ans += pn + 1
	}
	return ans
}
