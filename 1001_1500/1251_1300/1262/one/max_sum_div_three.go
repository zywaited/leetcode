package one

import "math"

func MaxSumDivThree(nums []int) int {
	// dp[i][j] 前i个数余数为j的最大值
	// 由于只跟前一个状态有关，所以只保留前一个
	// 为了减少后面的判断逻辑，余数为1和2的默认改为最小值
	// 这样
	// cn[1] = pn[1] + num
	// cn[2] = pn[2] + num
	// 两个表达式就不用管pn[1], pn[2]是否为0，
	// 不然会出现num整除，cn[1], cn[2] 值为num，整整除，与实际不符合
	pn := []int{0, math.MinInt32, math.MinInt32}
	cn := make([]int, 3)
	for _, num := range nums {
		switch num % 3 {
		case 0:
			cn[0] = pn[0] + num
			cn[1] = pn[1] + num
			cn[2] = pn[2] + num
		case 1:
			cn[0] = max(pn[0], pn[2]+num)
			cn[1] = max(pn[1], pn[0]+num)
			cn[2] = max(pn[2], pn[1]+num)
		case 2:
			cn[0] = max(pn[0], pn[1]+num)
			cn[1] = max(pn[1], pn[2]+num)
			cn[2] = max(pn[2], pn[0]+num)
		}
		pn, cn = cn, pn
	}
	return pn[0]
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
