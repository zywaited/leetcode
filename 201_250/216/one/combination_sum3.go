package one

func CombinationSum3(k int, n int) [][]int {
	rs := make([][]int, 0)
	var bc func(int, int, []int)
	bc = func(cn, tn int, cns []int) {
		if tn == 0 {
			if len(cns) == k {
				rs = append(rs, append(make([]int, 0, k), cns...))
			}
			return
		}
		for ; cn <= 9; cn++ {
			if tn < cn || len(cns) == k {
				return
			}
			bc(cn+1, tn-cn, append(cns, cn))
		}
	}
	bc(1, n, make([]int, 0, k))
	return rs
}
