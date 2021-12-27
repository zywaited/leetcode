package one

import "sort"

func LongestDiverseString(a int, b int, c int) string {
	// 计算出每个字符最大数量
	// 每两个字符穿插其他字符
	nums := [][2]int{
		{'a', min(a, (b+c+1)<<1)},
		{'b', min(b, (a+c+1)<<1)},
		{'c', min(c, (a+b+1)<<1)},
	}
	ans := make([]byte, nums[0][1]+nums[1][1]+nums[2][1])
	for i := 0; i < len(ans); i++ {
		// 每次排序获取最大值
		sort.Slice(nums, func(i, j int) bool {
			return nums[i][1] > nums[j][1]
		})
		// 是否连续3个一样
		if i <= 1 || ans[i-1] != ans[i-2] || ans[i-1] != byte(nums[0][0]) {
			ans[i] = byte(nums[0][0])
			nums[0][1]--
			continue
		}
		// 已经连续两个一样的，这里得用第二的数据
		ans[i] = byte(nums[1][0])
		nums[1][1]--
	}
	return string(ans)
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
