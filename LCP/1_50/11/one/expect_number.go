package one

// 对于非重复数字，那么看到的是一样的
// 对于重复数字，每个数字同时被查到的概率是1/N，但是有N个数字，所以就是1
// 所以这道题就是求去重后的数字数量
func ExpectNumber(scores []int) int {
	ans := 0
	// 空间换时间
	exists := make(map[int]byte, len(scores))
	for _, score := range scores {
		if exists[score] == 0 {
			exists[score] = 1
			ans++
		}
	}
	return ans
}
