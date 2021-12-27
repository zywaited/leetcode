package one

func SubarrayBitwiseORs(arr []int) int {
	curr := make(map[int]byte)
	ans := make(map[int]byte)
	for _, num := range arr {
		tmp := make(map[int]byte)
		for mask := range curr {
			tmp[mask|num] = 1
			ans[mask|num] = 1
		}
		tmp[num] = 1
		ans[num] = 1
		curr = tmp
	}
	return len(ans)
}
