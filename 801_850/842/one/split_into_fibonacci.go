package one

import "math"

func SplitIntoFibonacci(s string) []int {
	return bc([]byte(s), nil)
}

func bc(s []byte, prev []int) []int {
	if len(s) == 0 {
		if len(prev) < 3 {
			return nil
		}
		return prev
	}
	n := 0
	for i := range s {
		if (math.MaxInt32/10 < n) || (math.MaxInt32/10 == n && math.MaxInt32%10 < int(s[i]-'0')) {
			// 超过大小
			return nil
		}
		n = n*10 + int(s[i]-'0')
		if len(prev) >= 2 {
			if prev[len(prev)-1]+prev[len(prev)-2] < n {
				// 超过和了
				return nil
			}
			if prev[len(prev)-1]+prev[len(prev)-2] > n {
				continue
			}
		}
		ans := bc(s[i+1:], append(prev, n))
		if len(ans) > 0 {
			return ans
		}
		if i == 0 && n == 0 {
			// 0开头
			return nil
		}
	}
	// 无数据
	return nil
}
