package one

type node struct {
	word     string
	children []*node
}

func newNode() *node {
	return &node{children: make([]*node, 26)}
}

type searchNode struct {
	index int
	cn    *node
}

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{root: newNode()}
}

func (wd *WordDictionary) AddWord(word string) {
	cn := wd.root
	for index := range word {
		if cn.children[word[index]-'a'] == nil {
			cn.children[word[index]-'a'] = newNode()
		}
		cn = cn.children[word[index]-'a']
	}
	cn.word = word
}

func (wd *WordDictionary) Search(word string) bool {
	var q []searchNode
	wd.fill(&q, wd.root, 0, word)
	for len(q) > 0 {
		n := q[len(q)-1]
		q = q[:len(q)-1]
		if n.index+1 == len(word) {
			if len(n.cn.word) > 0 {
				return true
			}
			continue
		}
		wd.fill(&q, n.cn, n.index+1, word)
	}
	return false
}

func (wd *WordDictionary) fill(q *[]searchNode, cn *node, index int, word string) {
	if word[index] != '.' {
		if cn.children[word[index]-'a'] != nil {
			*q = append(*q, searchNode{index: index, cn: cn.children[word[index]-'a']})
		}
		return
	}
	for _, nn := range cn.children {
		if nn != nil {
			*q = append(*q, searchNode{index: index, cn: nn})
		}
	}
}
