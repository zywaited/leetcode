package one

// 元音字母
var il = map[byte]int{
	'a': 1 << 0,
	'e': 1 << 1,
	'i': 1 << 2,
	'o': 1 << 3,
	'u': 1 << 4,
}

func FindTheLongestSubstring(s string) int {
	ans := 0
	// 当某一位的值为1时, 就要找到上一次为1的起始位置
	// 这样就能满足偶数, 5个元音字母，2中状态，一共1<<5种组合
	status := 0
	indexes := make([]int, 1<<5)
	for index := range indexes {
		indexes[index] = -1
	}
	// 0
	indexes[0] = 0
	for index := range s {
		status ^= il[s[index]]
		// 当前状态已经出现了
		if indexes[status] >= 0 {
			ans = max(ans, index-indexes[status]+1)
			continue
		}
		indexes[status] = index + 1
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
