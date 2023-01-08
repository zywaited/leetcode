package one

const bitLen = 15

type node struct {
	num      int
	children []*node
}

type tree struct {
	root *node
}

func (tr *tree) add(num int) {
	cn := tr.root
	for i := bitLen; i >= 0; i-- {
		ni := (num >> i) & 1
		if cn.children[ni] == nil {
			cn.children[ni] = &node{children: make([]*node, 2)}
		}
		cn = cn.children[ni]
		cn.num++
	}
}

func CountPairs(nums []int, low int, high int) int {
	count := 0
	tr := tree{root: &node{children: make([]*node, 2)}}
	for _, num := range nums {
		count += getMinLimitCount(&tr, num, high+1) - getMinLimitCount(&tr, num, low)
		tr.add(num)
	}
	return count
}

func getMinLimitCount(tr *tree, num int, limit int) int {
	count := 0
	cn := tr.root
	for i := bitLen; i >= 0 && cn != nil; i-- {
		ni := (num >> i) & 1
		li := (limit >> i) & 1
		if li == 1 && cn.children[ni] != nil {
			// 异或为0的都小
			count += cn.children[ni].num
		}
		cn = cn.children[ni^li]
	}
	return count
}
