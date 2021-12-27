package two

import "math"

/**
 * 动态规划
 */

func dp(n, time int, mp map[int]int) int {
	// 是否已经处理过
	if _, ok := mp[n]; ok {
		return mp[n]
	}
	for {
		if n == 1 {
			break
		}
		// 等于3只需要两次
		if n == 3 {
			time += 2
			break
		}

		time++
		if n&1 == 0 {
			time += dp(n>>1, 0, mp)
			break
		}

		var (
			add int
			sub int
		)
		// 先加1
		// todo 如果需要考虑32的边界值
		if n == math.MaxInt32 {
			time++
			add = dp((math.MaxInt32>>1)+1, 0, mp)
		} else {
			add = dp(n+1, 0, mp)
		}
		// 再减1
		sub = dp(n-1, 0, mp)
		if add < sub {
			time += add
			break
		}
		time += sub
		break
	}
	// 保存值
	mp[n] = time
	return time
}

func IntegerReplacement(n int) int {
	return dp(n, 0, make(map[int]int))
}
