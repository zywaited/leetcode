package one

// 字典树
type trie struct {
	root *node
}

type node struct {
	char     byte
	children []*node
	word     bool
}

func newTrie() *trie {
	return &trie{&node{children: make([]*node, 26)}}
}

func (t *trie) insert(word string) {
	cn := t.root
	for i := range word {
		if cn.children[word[i]-'a'] == nil {
			cn.children[word[i]-'a'] = &node{char: word[i], children: make([]*node, 26)}
		}
		cn = cn.children[word[i]-'a']
	}
	cn.word = true
}

func MaxRectangle(words []string) []string {
	// 每个长度的单词有多少个
	wls := make(map[int][]string, 0)
	maxWordLen := 0
	t := newTrie()
	for _, word := range words {
		t.insert(word)
		wls[len(word)] = append(wls[len(word)], word)
		if len(word) > maxWordLen {
			maxWordLen = len(word)
		}
	}

	ans := make([]string, 0)
	max := 0
	var bc func([]string, []string, []*node)
	bc = func(ws []string, ps []string, ns []*node) {
		if len(ps) > maxWordLen || len(ws[0])*maxWordLen <= max {
			// 长度超过树高度或者不会比现有的大
			return
		}
		nns := make([]*node, 0, len(ns))
		f := true
		n := (*node)(nil)
		for index := range ws {
			nns = nns[:0]
			f = true
			for wi := range ws[index] {
				if len(ns) == 0 {
					n = t.root.children[ws[index][wi]-'a']
				} else {
					n = ns[wi].children[ws[index][wi]-'a']
				}
				if n == nil {
					break
				}
				nns = append(nns, n)
				f = f && n.word
			}
			if len(nns) < len(ws[0]) {
				continue
			}
			ps = append(ps, ws[index])
			if f && len(ws[0])*len(ps) > max {
				ans = ans[:0]
				ans = append(ans, ps...)
				max = len(ws[0]) * len(ps)
			}
			bc(ws, ps, nns)
			ps = ps[:len(ps)-1]
		}
	}
	for _, ws := range wls {
		bc(ws, nil, nil)
	}
	return ans
}
