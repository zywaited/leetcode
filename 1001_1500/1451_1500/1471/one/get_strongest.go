package one

import "sort"

func GetStrongest(arr []int, k int) []int {
	// 先排序获取中位数
	sort.Ints(arr)
	m := arr[(len(arr)-1)>>1]
	ans := make([]int, 0, k)
	l, r := 0, len(arr)-1
	for n := 0; n < k; n++ {
		fa := abs(m - arr[l])
		sa := abs(arr[r] - m)
		if sa >= fa {
			ans = append(ans, arr[r])
			r--
			continue
		}
		ans = append(ans, arr[l])
		l++
	}
	return ans
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
