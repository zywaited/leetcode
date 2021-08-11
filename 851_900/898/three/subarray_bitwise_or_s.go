package three

import "sort"

func SubarrayBitwiseORs(arr []int) int {
	indexes := [31][2]int{}
	for i := 0; i < len(indexes); i++ {
		indexes[i][0] = i
	}
	ans := make(map[int]byte)
	for index, num := range arr {
		ans[num] = 1
		if num == 0 {
			continue
		}
		sort.Slice(indexes[:], func(i, j int) bool {
			return indexes[i][1] > indexes[j][1]
		})
		mask := num
		for i := 0; i < len(indexes); i++ {
			if indexes[i][1] > 0 {
				mask |= arr[indexes[i][1]-1]
				ans[mask] = 1
			}
			if num&(1<<uint(indexes[i][0])) > 0 {
				indexes[i][1] = index + 1
			}
		}
	}
	return len(ans)
}
