package two

import "sort"

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	indexNums := [][]int{}
	for index, num := range nums {
		indexNums = append(indexNums, []int{num, index})
	}
	sort.Slice(indexNums, func(i, j int) bool {
		return indexNums[i][0] < indexNums[j][0]
	})
	ans := make([]bool, len(l))
	for i := 0; i < len(ans); i++ {
		if r[i]-l[i] == 1 {
			ans[i] = true
			continue
		}
		pn := 0
		diff := 0
		n := 0
		ni := 0
		for ; ni < len(indexNums); ni++ {
			if indexNums[ni][1] < l[i] || indexNums[ni][1] > r[i] {
				continue
			}
			switch n {
			case 0:
				n++
			case 1:
				diff = indexNums[ni][0] - pn
				n++
			case 2:
				if indexNums[ni][0]-pn != diff {
					n++
				}
			}
			if n > 2 {
				break
			}
			pn = indexNums[ni][0]
		}
		ans[i] = ni == len(indexNums)
	}
	return ans
}
