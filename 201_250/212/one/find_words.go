package one

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
	cn := t.root
	for i := range word {
		if cn.children[word[i]-'a'] == nil {
			cn.children[word[i]-'a'] = &node{val: word[i], children: make([]*node, 26)}
		}
		cn = cn.children[word[i]-'a']
	}
	cn.word = true
}

/**
 * 回溯结束信号
 * 1 列或者行溢出
 * 2 当前点已经在回溯中(四周扩散会回溯回来)
 * 3 已经找到了所有需要的word
 */
func bc(t *Trie, board [][]byte, row, col int, bs []byte, num int, find map[string]bool, cn *node, box [][]bool) {
	// 判断是否已经到头
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) || box[row][col] || len(find) == num {
		return
	}
	// 查找是否存在
	cn = cn.children[board[row][col]-'a']
	if cn == nil {
		return
	}
	bs = append(bs, board[row][col])
	if cn.word {
		find[string(bs)] = true
	}
	// 向四周扩散
	box[row][col] = true
	// 优化点：每次都带上cn, 下一次查找就从当前继续
	bc(t, board, row, col-1, bs, num, find, cn, box)
	bc(t, board, row, col+1, bs, num, find, cn, box)
	bc(t, board, row-1, col, bs, num, find, cn, box)
	bc(t, board, row+1, col, bs, num, find, cn, box)
	box[row][col] = false
}

func FindWords(board [][]byte, words []string) []string {
	var res []string
	row := len(board)
	if len(board) < 1 || len(board[0]) < 1 {
		return res
	}
	col := len(board[0])
	// 构建前缀树
	t := newTrie()
	for _, word := range words {
		t.Insert(word)
	}
	// 初始化box
	box := make([][]bool, len(board))
	for i := range box {
		box[i] = make([]bool, len(board[0]))
	}
	find := make(map[string]bool)
	num := len(words)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			bc(t, board, i, j, nil, num, find, t.root, box)
		}
	}
	for s := range find {
		res = append(res, s)
	}
	return res
}
