package one

import (
	"math"
	"sort"
)

func MakeArrayIncreasing(arr1 []int, arr2 []int) int {
	if len(arr1) <= 1 {
		return 0
	}

	// 先对arr2去重排序
	ex := make(map[int]byte, len(arr2))
	arr3 := make([]int, 0, len(arr2))
	for _, num := range arr2 {
		if ex[num] == 1 {
			continue
		}
		ex[num] = 1
		arr3 = append(arr3, num)
	}
	sort.Ints(arr3)
	search := func(s int) (int, bool) {
		// 二分查找最近一个比s大的数
		l, r := 0, len(arr3)-1
		m := 0
		i := -1
		for l <= r {
			m = l + (r-l)>>1
			if arr3[m] == s {
				// arr3没有重复数字
				// 则下一个一定比s大
				i = m + 1
				break
			}
			if arr3[m] < s {
				l = m + 1
				continue
			}
			// 比s大，但还要继续往左边找
			i = m
			r = m - 1
		}
		if i > -1 && i < len(arr3) {
			return arr3[i], true
		}
		return -1, false
	}

	// i代表修改次数
	// i次修改后是否有效
	inValid := make(map[int]byte)
	// j代表arr1的前j数
	// dp[i][j]当前的右边界(边界越小越好)
	dp := make([][]int, len(arr1)+1)
	dp[0] = make([]int, len(arr1)+1)
	i := 0
	for i = range dp[0] {
		dp[0][i] = math.MaxInt32
	}
	dp[0][0] = math.MinInt32
	m := 0     // 替换元素
	f := false // 是否找到替换元素
	d := false // 是否递增
	j := 1
	for ; j < len(dp); j++ {
		// 从头开始计算最小值
		for i = 0; i <= j; i++ {
			if inValid[i] == 1 {
				continue
			}
			if len(dp[i]) == 0 {
				dp[i] = make([]int, len(arr1)+1)
				for k := range dp[i] {
					dp[i][k] = math.MaxInt32
				}
			}
			d = false
			if arr1[j-1] > dp[i][j-1] {
				// 满足递增，直接更新边界
				dp[i][j] = arr1[j-1]
				d = true
			}
			// 更改dp[i-1][j-1]的边界值也是dp[i][j]
			// 所以比较两个大小
			if i > 0 && dp[i-1][j-1] != math.MaxInt32 {
				m, f = search(dp[i-1][j-1])
				if f && m < dp[i][j] {
					dp[i][j] = m
					d = true
				}
			}
			if !d {
				inValid[i] = 1
				continue
			}
			// 是否满足了
			if j == len(dp)-1 {
				return i
			}
		}
	}
	return -1
}
