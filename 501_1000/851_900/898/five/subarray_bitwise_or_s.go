package five

func SubarrayBitwiseORs(arr []int) int {
	indexes := [31][]int{}
	for i := 0; i < len(indexes); i++ {
		indexes[i] = []int{i, 0}
	}
	swap := [31][]int{}
	ans := make(map[int]byte)
	for index, num := range arr {
		ans[num] = 1
		if num == 0 {
			continue
		}
		mask := num
		swapIndex := 0
		for i := 0; i < len(indexes); i++ {
			if indexes[i][1] > 0 {
				mask |= arr[indexes[i][1]-1]
				ans[mask] = 1
			}
			if num&(1<<uint(indexes[i][0])) > 0 {
				indexes[i][1] = index + 1
				swap[swapIndex] = indexes[i]
				swapIndex++
			}
		}
		for i := 0; i < len(indexes); i++ {
			if num&(1<<uint(indexes[i][0])) == 0 {
				swap[swapIndex] = indexes[i]
				swapIndex++
			}
		}
		indexes, swap = swap, indexes
	}
	return len(ans)
}
