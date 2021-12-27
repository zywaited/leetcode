package one

func PalindromePairs(words []string) [][]int {
	indexes := make(map[int]map[int]bool, len(words))
	wm := make(map[string]int)
	for index, word := range words {
		wm[word] = index + 1
		indexes[index] = make(map[int]bool)
	}
	ans := make([][]int, 0)
	for index, word := range words {
		if len(word) == 0 {
			// 这种情况会被其他word处理
			continue
		}
		// 反转字符串
		rw := revert(word)
		dp := make([][]bool, len(word)) // 计算区间是否是回文串
		for i := 0; i < len(word); i++ {
			dp[i] = make([]bool, len(word))
			for j := i; j >= 0; j-- {
				dp[j][i] = word[j] == word[i] && (i-j <= 2 || dp[j+1][i-1])
				dp[i][j] = dp[j][i] // 反向也赋值
			}
		}
		for ci := 0; ci < len(word); ci++ {
			if ci == 0 {
				nextIndex := wm[rw] - 1
				if nextIndex >= 0 && nextIndex != index && !indexes[index][nextIndex] {
					ans = append(ans, []int{index, nextIndex})
					indexes[index][nextIndex] = true
				}
				// 本身就是一个回文串
				if dp[0][len(word)-1] && wm[""] > 0 {
					ans = append(ans, []int{index, wm[""] - 1})
					ans = append(ans, []int{wm[""] - 1, index})
				}
				continue
			}
			// 左边
			if dp[0][ci-1] {
				nextIndex := wm[rw[0:len(word)-ci]] - 1
				if nextIndex >= 0 && nextIndex != index && !indexes[nextIndex][index] {
					ans = append(ans, []int{nextIndex, index})
					indexes[nextIndex][index] = true
				}
			}
			// 右边
			if dp[len(word)-ci][len(word)-1] {
				nextIndex := wm[rw[ci:]] - 1
				if nextIndex >= 0 && nextIndex != index && !indexes[index][nextIndex] {
					ans = append(ans, []int{index, nextIndex})
					indexes[index][nextIndex] = true
				}
			}
		}
	}
	return ans
}

func revert(s string) string {
	rs := make([]byte, 0, len(s))
	for index := len(s) - 1; index >= 0; index-- {
		rs = append(rs, s[index])
	}
	return string(rs)
}
