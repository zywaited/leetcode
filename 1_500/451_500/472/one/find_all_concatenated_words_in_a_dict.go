package one

import "sort"

type node struct {
	word     bool
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
			n.children[word[index]-'a'] = &node{children: make([]*node, 26)}
		}
		n = n.children[word[index]-'a']
	}
	n.word = true
}

func FindAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ans := make([]string, 0, len(words))
	dt := newDictTree()
	var dfs func(word string) bool
	dfs = func(word string) bool {
		if word == "" {
			return true
		}

		cn := dt.root
		for i := range word {
			if cn.children == nil || cn.children[word[i]-'a'] == nil {
				return false
			}
			cn = cn.children[word[i]-'a']
			if cn.word && dfs(word[i+1:]) {
				return true
			}
		}
		return false
	}
	for _, word := range words {
		if word == "" {
			continue
		}
		if dfs(word) {
			ans = append(ans, word)
			continue
		}
		dt.insert(word)
	}
	return ans
}
