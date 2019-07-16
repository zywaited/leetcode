package one

type Trie struct {
	root *node
}

type node struct {
	val byte

	children []*node

	word bool
}

func Constructor() *Trie {
	return &Trie{&node{children: make([]*node, 26)}}
}

func (t *Trie) Insert(word string) {
	cn := t.root
	for i := range word {
		if cn.children[word[i]-'a'] == nil {
			cn.children[word[i]-'a'] = &node{val: word[i], children: make([]*node, 26)}
		}
		cn = cn.children[word[i]-'a']
	}
	cn.word = true
}

func (t *Trie) Search(word string) bool {
	cn := t.root
	for i := range word {
		if cn.children[word[i]-'a'] == nil {
			return false
		}
		cn = cn.children[word[i]-'a']
	}
	return cn.word
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) < 1 {
		return false
	}
	cn := t.root
	for i := range prefix {
		if cn.children[prefix[i]-'a'] == nil {
			return false
		}
		cn = cn.children[prefix[i]-'a']
	}
	return true
}
