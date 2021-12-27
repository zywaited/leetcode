package one

func FindNumOfValidWords(words []string, puzzles []string) []int {
	rs := make([]int, len(puzzles))
	wn := make(map[int]int) // 同种单词的数量(有相同字母)
	for _, word := range words {
		score := 0
		for i := range word {
			score |= 1 << (word[i] - 'a')
		}
		wn[score]++
	}
	// 因为第一个字符必须存在，所以减少一个（题目说明长度为7，所以不用len(puzzle)计算了）
	// 写在这里而不是for中是不用每次都初始化scores，节约内存分配
	cn := 1 << 6
	scores := make([]int, cn)
	for i, puzzle := range puzzles {
		if len(puzzle) == 0 {
			continue
		}
		// 循环wn数组应该会超时
		// 不过puzzle的长度为7，并且第一个字符必须在word中，这样数量少
		// 先获取所有的组合
		for j := 0; j < cn; j++ {
			scores[j] = 1 << (puzzle[0] - 'a')
			for k := 1; k < len(puzzle); k++ {
				if j&(1<<(uint(k-1))) != 0 {
					scores[j] |= 1 << (puzzle[k] - 'a')
				}
			}
		}
		// 判断是否存在谜底
		for _, score := range scores {
			rs[i] += wn[score]
		}
	}
	return rs
}
