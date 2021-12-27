package one

import "sort"

func MaxJumps(arr []int, d int) int {
	// 先按照值升序排序索引 (先计算小的，到一步一步推导到高的, arr[i] > arr[k])
	indexes := make([]int, 0, len(arr))
	for index := range arr {
		indexes = append(indexes, index)
	}
	sort.Slice(indexes, func(i, j int) bool {
		return arr[indexes[i]] < arr[indexes[j]]
	})
	// 储存当前索引得最大访问个数
	dp := make([]int, len(arr))
	ans := 0
	for _, index := range indexes {
		dp[index] = 1
		// i - x >= 0 && 0 < x <= d && arr[i] > arr[k]
		for i := index - 1; i >= 0 && i >= index-d; i-- {
			if arr[i] >= arr[index] {
				break
			}
			dp[index] = max(dp[index], dp[i]+1)
		}
		// i + x < arr.length && 0 < x <= d && arr[i] > arr[k]
		for i := index + 1; i < len(arr) && i <= index+d; i++ {
			if arr[i] >= arr[index] {
				break
			}
			dp[index] = max(dp[index], dp[i]+1)
		}
		ans = max(ans, dp[index])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
