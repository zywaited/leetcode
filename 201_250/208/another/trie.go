package another

type Trie struct {
	root *node
}

type node struct {
	prefix string

	children []*node

	// 为了不使用递归回溯
	parent *node

	// 原始数据
	s string
}

func (n *node) findChild(label byte) *node {
	for _, child := range n.children {
		if child.prefix[0] == label {
			return child
		}
	}
	return nil
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}

/** Initialize your data structure here. */
func Constructor() *Trie {
	return &Trie{root: &node{}}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	if len(word) < 1 {
		return
	}
	// 找到能匹配的节点
	cn := t.root.findChild(word[0])
	if cn == nil {
		// 只能放在root下
		t.root.addChild(&node{prefix: word, parent: cn, s: word})
		return
	}
	search := word
	for {
		sl := len(search)
		pl := len(cn.prefix)
		l := 0

		// 计算可以匹配的长度
		max := sl
		if pl < max {
			max = pl
		}
		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}

		// 未匹配
		if l == 0 {
			// 当前节点的子节点
			cn.parent.addChild(&node{prefix: search, parent: cn, s: word})
			return
		}
		// 匹配一部分
		// 被所有的节点
		if l < pl {
			// 拆分节点
			n := &node{prefix: cn.prefix[l:], parent: cn, children: cn.children, s: cn.s}
			// 重置父节点
			cn.prefix = search[:l]
			cn.children = nil
			cn.s = ""
			cn.addChild(n)
			// 如果匹配了搜索
			if l == sl {
				cn.s = word
				return
			}
			// 新的子节点
			cn.addChild(&node{prefix: search[l:], parent: cn, s: word})
			return
		}
		// 搜索
		if l < sl {
			// 找到下一个子节点
			search = search[l:]
			tmp := cn.findChild(search[0])
			if tmp != nil {
				cn = tmp
				continue
			}
			// 不存在则创建子节点
			cn.addChild(&node{prefix: search, parent: cn, s: word})
			return
		}
		// 全匹配
		if len(cn.s) < 1 {
			// 只是分割出来的，没有数据
			cn.s = word
		}
		return
	}
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	if len(word) < 1 {
		return false
	}
	cn := t.root.findChild(word[0])
	if cn == nil {
		return false
	}
	search := word
	for {
		sl := len(search)
		pl := len(cn.prefix)
		max := sl
		l := 0
		if pl < max {
			max = pl
		}
		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}
		// 匹配失败
		if l == 0 || l < pl {
			return false
		}
		if l < sl {
			// 继续子节点匹配
			search = search[l:]
			tmp := cn.findChild(search[0])
			if tmp != nil {
				cn = tmp
				continue
			}
			return false
		}
		// 验证是否真的
		return len(cn.s) == len(word)
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	pl := len(prefix)
	if pl < 1 {
		return false
	}
	cn := t.root.findChild(prefix[0])
	if cn == nil {
		return false
	}
	search := prefix
	for {
		sl := len(search)
		pl := len(cn.prefix)
		max := sl
		l := 0
		if pl < max {
			max = pl
		}
		for ; l < max && search[l] == cn.prefix[l]; l++ {
		}
		// 匹配失败
		if l == 0 || (l < pl && l < sl) {
			return false
		}
		if l < sl {
			// 继续子节点匹配
			search = search[l:]
			tmp := cn.findChild(search[0])
			if tmp != nil {
				cn = tmp
				continue
			}
			return false
		}
		return true
	}
}
