package one

// 回溯
func Combine(n int, k int) [][]int {
	rs := make([][]int, 0)
	var bc func(int, []int)
	bc = func(fn int, cns []int) {
		if len(cns) == k {
			rs = append(rs, append(make([]int, 0, k), cns...))
			return
		}
		for sn := fn; sn <= n; sn++ {
			bc(sn+1, append(cns, sn))
		}
	}
	bc(1, make([]int, 0, k))
	return rs
}
