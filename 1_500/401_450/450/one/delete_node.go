package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	cn := root
	parent := (*TreeNode)(nil)
	for cn != nil && cn.Val != key {
		parent = cn
		if key < cn.Val {
			cn = cn.Left
			continue
		}
		cn = cn.Right
	}
	if cn == nil {
		return root
	}
	if cn.Right == nil {
		if parent == nil {
			// ans := cn.Left
			// cn.Left = nil
			// return ans
			return cn.Left
		}
		if parent.Left == cn {
			parent.Left = cn.Left
			// cn.Left = nil
			return root
		}
		parent.Right = cn.Left
		// cn.Right = nil
		return root
	}
	parent = cn
	sucessor := cn.Right
	for sucessor != nil && sucessor.Left != nil {
		parent = sucessor
		sucessor = sucessor.Left
	}
	// note: 这里应该换成交换节点
	cn.Val = sucessor.Val
	// right := sucessor.Right
	// sucessor.Right = nil
	if parent.Left == sucessor {
		parent.Left = sucessor.Right
		return root
	}
	parent.Right = sucessor.Right
	return root
}
