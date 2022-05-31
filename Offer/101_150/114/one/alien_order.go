package one

func AlienOrder(words []string) string {
	priorities := [26][]int{}
	for i := 0; i < len(words); i++ {
		f := false
		for ci := 0; ci < len(words[i]); ci++ {
			// 找到优先级
			if priorities[words[i][ci]-'a'] == nil {
				priorities[words[i][ci]-'a'] = []int{}
			}
			if i == 0 || len(words[i-1]) <= ci {
				continue
			}
			// 找到优先级
			if !f && words[i-1][ci] != words[i][ci] {
				priorities[words[i][ci]-'a'] = append(priorities[words[i][ci]-'a'], int(words[i-1][ci]-'a'))
				f = true
			}
		}
		if !f && i > 0 && len(words[i]) < len(words[i-1]) {
			return ""
		}
	}
	// 拓扑排序
	bs := make([]byte, 0, 26)
	visited := [26]int{}
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if visited[i] == -1 {
			return true
		}
		// 循环依赖
		if visited[i] == 1 {
			return false
		}
		visited[i] = 1
		for _, pi := range priorities[i] {
			if !dfs(pi) {
				return false
			}
		}
		visited[i] = -1
		bs = append(bs, byte(i)+'a')
		return true
	}
	for i := range priorities {
		// 不存在当前字母
		if priorities[i] == nil {
			continue
		}
		if !dfs(i) {
			return ""
		}
	}
	return string(bs)
}
