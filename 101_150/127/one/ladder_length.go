package one

func LadderLength(beginWord string, endWord string, wordList []string) int {
	type node struct {
		// 当前单词
		word string
		// 深度
		level int
	}
	var (
		// 存储预处理数据
		// abc可以预处理成*bc,a*c,ab*
		mp = make(map[string][]string)
		// 下一深度的字符串
		queue []node
		// 已经处理过的单词，下一层访问已经处理过的单词直接跳过，因为深度比上一层大
		used = make(map[string]byte)
	)
	// 中间数据
	var (
		word  string
		next  string
		level int
		// 通配符
		wildcard string
	)
	// 预处理
	exist := false
	for _, word = range wordList {
		if word == beginWord {
			continue
		}
		if word == endWord {
			exist = true
		}
		for i := 0; i < len(word); i++ {
			wildcard = word[:i] + "*" + word[i+1:]
			mp[wildcard] = append(mp[wildcard], word)
		}
	}
	if !exist {
		return 0
	}
	// 转换
	queue = append(queue, node{word: beginWord, level: 1})
	for len(queue) != 0 {
		word, level = queue[0].word, queue[0].level
		queue = queue[1:]
		for i := 0; i < len(word); i++ {
			wildcard = word[:i] + "*" + word[i+1:]
			// 不存在
			if mp[wildcard] == nil {
				continue
			}
			for _, next = range mp[wildcard] {
				// 是否找到
				if next == endWord {
					return level + 1
				}
				// 是否已经用过
				if used[next] > 0 {
					continue
				}
				// 加入队列中
				used[next] = 1
				queue = append(queue, node{word: next, level: level + 1})
			}
		}
	}
	// 找不到
	return 0
}
