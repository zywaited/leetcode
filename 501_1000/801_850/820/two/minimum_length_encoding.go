package two

func MinimumLengthEncoding(words []string) int {
	t := newTrie()
	t.Insert(words[0])
	ans := len(words[0]) + 1
	for i := 1; i < len(words); i++ {
		cn, l := t.Find(words[i])
		if l == len(words[i]) {
			// word被包含了
			continue
		}
		if l > 0 && cn.Word() {
			// cn被包含了
			ans -= l + 1
			// 去除单词标记
			cn.Reset()
		}
		ans += len(words[i]) + 1
		if cn != nil {
			cn.Insert(words[i][:len(words[i])-l])
		}
	}
	return ans
}

type Trie struct {
	root *node
}

type node struct {
	val byte

	children []*node

	word bool
}

func newTrie() *Trie {
	return &Trie{&node{children: make([]*node, 26)}}
}

func (t *Trie) Insert(word string) {
	t.root.Insert(word)
}

func (t *Trie) Find(word string) (*node, int) {
	cn := t.root
	for i := len(word) - 1; i >= 0; i-- {
		if cn.children[word[i]-'a'] != nil {
			cn = cn.children[word[i]-'a']
			continue
		}
		return cn, len(word) - i - 1
	}
	return cn, len(word)
}

func (n *node) Insert(word string) {
	cn := n
	for i := len(word) - 1; i >= 0; i-- {
		if cn.children[word[i]-'a'] == nil {
			cn.children[word[i]-'a'] = &node{val: word[i], children: make([]*node, 26)}
		}
		cn = cn.children[word[i]-'a']
	}
	cn.word = true
}

func (n *node) Word() bool {
	return n.word
}

func (n *node) Reset() {
	n.word = false
}
