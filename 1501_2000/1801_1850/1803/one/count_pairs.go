package one

type node struct {
	index    int
	num      int
	children []*node
}

type tree struct {
	root *node
}

func (tr *tree) add(num int) {
	cn := tr.root
	for i := cn.index - 1; i >= 0; i-- {
		ni := (num >> i) & 1
		if cn.children[ni] == nil {
			cn.children[ni] = &node{index: i, children: make([]*node, 2)}
		}
		cn = cn.children[ni]
		cn.num++
	}
}

func CountPairs(nums []int, low int, high int) int {
	count := 0
	tr := tree{root: &node{index: 15, children: make([]*node, 2)}}
	for _, num := range nums {
		count += getCount(&tr, num, low, high+1)
		tr.add(num)
	}
	return count
}

func getCount(tr *tree, num int, low int, high int) int {
	count := 0
	cn := tr.root
	diff := false
	for cn != nil && cn.index > 0 {
		ni := (num >> (cn.index - 1)) & 1
		hi := (high >> (cn.index - 1)) & 1
		if hi == 0 {
			// 只能选择一样的位
			cn = cn.children[ni]
			continue
		}
		li := (low >> (cn.index - 1)) & 1
		if diff && cn.children[ni] != nil {
			// diff == true 一定比low大
			count += cn.children[ni].num
		}
		if li != hi && !diff {
			// 与high前缀一致，此时异或为0一定比high小, 后面位选择比low大的
			count += getLowCount(cn.children[ni], num, low-1)
			diff = true
		}
		cn = cn.children[ni^hi]
	}
	return count
}

func getLowCount(cn *node, num int, low int) int {
	count := 0
	for cn != nil && cn.index > 0 {
		ni := (num >> (cn.index - 1)) & 1
		li := (low >> (cn.index - 1)) & 1
		if li == 1 {
			// 只能选择异或为1的位
			cn = cn.children[ni^li]
			continue
		}
		if cn.children[ni^1] != nil {
			// 异或为1一定比low大
			count += cn.children[ni^1].num
		}
		cn = cn.children[ni]
	}
	return count
}
