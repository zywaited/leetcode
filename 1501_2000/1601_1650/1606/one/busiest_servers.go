package one

func BusiestServers(k int, arrival []int, load []int) []int {
	tr := &lineTree{pairs: make([]*pair, k), root: &lineTreeNode{low: 0, high: k - 1}}
	for i := range arrival {
		tr.update(i%k, arrival[i], arrival[i]+load[i])
	}
	ans := make([]int, 0, k)
	for _, p := range tr.pairs {
		if p == nil {
			continue
		}
		if len(ans) > 0 && p.times < tr.pairs[ans[0]].times {
			continue
		}
		if len(ans) > 0 && tr.pairs[ans[0]].times < p.times {
			ans = ans[:0]
		}
		ans = append(ans, p.index)
	}
	return ans
}

type lineTree struct {
	pairs []*pair
	root  *lineTreeNode
}

func (tree *lineTree) update(index, curr, next int) {
	if tree.pairs[index] == nil {
		tree.pairs[index] = &pair{index: index}
	}
	p := tree.pairs[index]
	if curr < p.valid {
		p = tree.root.get(index, len(tree.pairs)-1, curr)
		if p == nil && index-1 >= 0 {
			p = tree.root.get(0, index-1, curr)
		}
	}
	if p == nil {
		return
	}
	p.times++
	p.valid = next
	tree.root.update(p)
}

type lineTreeNode struct {
	low   int
	high  int
	min   *pair
	left  *lineTreeNode
	right *lineTreeNode
}

func (node *lineTreeNode) update(server *pair) {
	if server.index < node.low || node.high < server.index {
		return
	}
	node.initLR()
	if node.low == node.high {
		node.min = server
		return
	}
	node.left.update(server)
	node.right.update(server)
	node.min = min(node.left.min, node.right.min)
}

func (node *lineTreeNode) initLR() {
	if node.low == node.high {
		return
	}
	if node.left == nil {
		node.left = &lineTreeNode{low: node.low, high: node.low + (node.high-node.low)>>1}
	}
	if node.right == nil {
		node.right = &lineTreeNode{low: node.left.high + 1, high: node.high}
	}
}

func (node *lineTreeNode) get(low, high, curr int) *pair {
	if node.high < low || high < node.low {
		return nil
	}
	if low <= node.low && node.high <= high {
		if curr < node.min.valid {
			return nil
		}
		if node.low == node.high {
			return node.min
		}
	}
	node.initLR()
	var p *pair
	p = node.left.get(low, high, curr)
	if p == nil {
		p = node.right.get(low, high, curr)
	}
	return p
}

type pair struct {
	index int
	times int
	valid int
}

func min(f, s *pair) *pair {
	if f == nil {
		return s
	}
	if s == nil {
		return f
	}
	if f.valid < s.valid || (f.valid == s.valid && f.index < s.index) {
		return f
	}
	return s
}
