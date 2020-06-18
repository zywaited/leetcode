package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	cn *TreeNode
	d  int // 深度
	l  int // 占用长度
}

func RecoverFromPreorder(S string) *TreeNode {
	// 先取出根节点，题目明确有一个节点
	index := 0
	n := newNode(index, S)
	root := n.cn
	stack := []*node{n}
	sl := len(stack)
	index += n.l
	for index < len(S) {
		n = newNode(index, S)
		for stack[sl-1].d >= n.d {
			sl--
		}
		if stack[sl-1].cn.Left == nil {
			// 左节点
			stack[sl-1].cn.Left = n.cn
		} else {
			// 右节点
			stack[sl-1].cn.Right = n.cn
			sl--
		}
		stack = stack[:sl]
		stack = append(stack, n)
		sl++
		index += n.l
	}
	return root
}

// 获取节点深度、节点值和下次索引
func newNode(index int, S string) *node {
	n := &node{cn: &TreeNode{}}
	for ; index < len(S); index++ {
		if S[index] >= '0' && S[index] <= '9' {
			n.cn.Val = n.cn.Val*10 + int(S[index]-'0')
			n.l++
			continue
		}
		if n.cn.Val > 0 {
			// 下一个节点了
			break
		}
		n.l++
		n.d++
	}
	return n
}
