package two

func MultiSearch(big string, smalls []string) [][]int {
	t := newTree()
	for index, small := range smalls {
		t.insert(index, small)
	}
	ans := make([][]int, len(smalls))
	for bi := 0; bi < len(big); bi++ {
		t.fill(ans, big, bi)
	}
	return ans
}

// 字典树
type tree struct {
	root []*node
}

type node struct {
	index    int // smalls index
	children []*node
}

func newTree() *tree {
	return &tree{root: make([]*node, 26)}
}

func newNode() *node {
	return &node{index: -1, children: make([]*node, 26)}
}

func (t *tree) insert(index int, small string) {
	if len(small) == 0 {
		return
	}
	cn := t.root
	ni := byte(0)
	for i := range small {
		ni = small[i] - 'a'
		if cn[ni] == nil {
			cn[ni] = newNode()
		}
		if i == len(small)-1 {
			cn[ni].index = index
			break
		}
		cn = cn[ni].children
	}
	return
}

func (t *tree) fill(ans [][]int, big string, bi int) {
	cn := t.root
	ni := byte(0)
	for index := bi; index < len(big); index++ {
		ni = big[index] - 'a'
		if len(cn) == 0 || cn[ni] == nil {
			return
		}
		if cn[ni].index > -1 {
			ans[cn[ni].index] = append(ans[cn[ni].index], bi)
		}
		cn = cn[ni].children
	}
}
