package two

// O(N+NLogN)
func MinSubArrayLen(s int, nums []int) int {
	// 前缀和数组O(N)
	ps := make([]int, len(nums)+1)
	for i, num := range nums {
		ps[i+1] = ps[i] + num
	}
	// 二分搜索满足的长度O(LogN)
	bs := func(ci int) int {
		l := len(nums) + 1
		if ps[ci] < s {
			return l
		}
		si, ei := 0, ci
		for si <= ei {
			mi := si + (ei-si)>>1
			if ps[ci]-ps[mi] < s {
				ei = mi - 1
				continue
			}
			si = mi + 1
			l = min(l, ci-mi)
		}
		return l
	}
	// O(NLogN)
	ans := len(nums) + 1
	for i := 1; i <= len(nums); i++ {
		ans = min(ans, bs(i))
	}
	if ans > len(nums) {
		ans = 0
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
