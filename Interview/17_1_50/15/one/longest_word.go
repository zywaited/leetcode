package one

type node struct {
	char     byte
	word     int
	children []*node
}

type dictTree struct {
	root *node
}

func newDictTree() *dictTree {
	return &dictTree{root: &node{children: make([]*node, 26)}}
}

func (dt *dictTree) insert(word string) {
	n := dt.root
	for index := range word {
		if n.children[word[index]-'a'] == nil {
			n.children[word[index]-'a'] = &node{char: word[index], children: make([]*node, 26)}
		}
		n = n.children[word[index]-'a']
	}
	n.word++
}

func LongestWord(words []string) string {
	dt := newDictTree()
	for _, word := range words {
		dt.insert(word)
	}
	var bc func(int, int, *node) bool
	bc = func(i, s int, n *node) bool {
		// 这里其实可以记录中间状态
		// words[i][s:]
		if s == len(words[i]) {
			return true
		}
		for e := s; e < len(words[i]); e++ {
			n = n.children[words[i][e]-'a']
			if n == nil {
				return false
			}
			if n.word == 0 {
				continue
			}
			if !bc(i, e+1, dt.root) {
				continue
			}
			if s == 0 && e == len(words[i])-1 && n.word == 1 {
				return false
			}
			return true
		}
		return false
	}
	ans := ""
	for i := range words {
		if !bc(i, 0, dt.root) || len(words[i]) < len(ans) {
			continue
		}
		if len(words[i]) > len(ans) || words[i] < ans {
			ans = words[i]
		}
	}
	return ans
}
