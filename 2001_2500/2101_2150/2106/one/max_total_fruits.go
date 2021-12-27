package one

import "sort"

func MaxTotalFruits(fruits [][]int, startPos int, k int) int {
	si := sort.Search(len(fruits), func(i int) bool {
		return startPos <= fruits[i][0]
	})
	// 先取右边的数量
	curr := 0
	i := si
	for ; i < len(fruits) && fruits[i][0]-startPos <= k; i++ {
		curr += fruits[i][1]
	}
	ri := i - 1
	rk := 0
	if si <= ri {
		rk = fruits[ri][0] - startPos
	}
	ans := curr
	for li := si - 1; li >= 0; li-- {
		for ; (k < rk*2+startPos-fruits[li][0] && k < (startPos-fruits[li][0])*2+rk) && si <= ri; ri-- {
			curr -= fruits[ri][1]
			if si < ri {
				rk -= fruits[ri][0] - fruits[ri-1][0]
				continue
			}
			rk -= fruits[ri][0] - startPos
		}
		if (rk*2+startPos-fruits[li][0] <= k) || ((startPos-fruits[li][0])*2+rk <= k) {
			curr += fruits[li][1]
			ans = max(ans, curr)
			continue
		}
		break
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
