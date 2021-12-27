package one

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	type node struct {
		// 当前单词
		word string
		// 转换路径
		path []string
		// 深度
		level int
	}
	var (
		// 存储预处理数据
		// abc可以预处理成*bc,a*c,ab*
		mp = make(map[string][]string)
		// 已经处理过的单词，下一层访问已经处理过的单词直接跳过，因为深度比上一层大
		used = make(map[string]int)
		// 下一深度的字符串
		queue []node
		rs    [][]string
	)
	// 预处理
	exist := false
	for _, word := range wordList {
		if word == beginWord {
			continue
		}
		if word == endWord {
			exist = true
		}
		for i := 0; i < len(word); i++ {
			wildcard := word[:i] + "*" + word[i+1:]
			mp[wildcard] = append(mp[wildcard], word)
		}
	}
	if !exist {
		return rs
	}
	// 转换
	queue = append(queue, node{word: beginWord, path: []string{beginWord}, level: 1})
	for len(queue) != 0 {
		word, path, level := queue[0].word, queue[0].path, queue[0].level
		// 可能已经有脏数据加入队列中
		if len(rs) > 0 && len(rs[0])-1 < len(path) {
			break
		}
		queue = queue[1:]
		for i := 0; i < len(word); i++ {
			wildcard := word[:i] + "*" + word[i+1:]
			// 不存在
			if mp[wildcard] == nil {
				continue
			}
			for _, next := range mp[wildcard] {
				// 是否找到
				if next == endWord {
					// 为了节约一点点内存
					// append(path, next)
					r := make([]string, len(path)+1)
					copy(r, path)
					r[len(r)-1] = next
					rs = append(rs, r)
					continue
				}
				// 上级已经使用
				if used[next] > 0 && used[next] <= level {
					continue
				}
				// 如果已经找到了
				if len(rs) > 0 && len(rs[0])-1 < len(path) {
					continue
				}
				// 加入队列
				used[next] = level + 1
				// 为了节约一点点内存
				// append(append([]string{}, path...), next)
				n := node{word: next, path: make([]string, len(path)+1), level: level + 1}
				copy(n.path, path)
				n.path[len(n.path)-1] = next
				queue = append(queue, n)
			}
		}
	}
	// 找不到
	return rs
}
