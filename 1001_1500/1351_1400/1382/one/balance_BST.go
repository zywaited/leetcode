package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历数据，这样数据就是从小到大排列
// 然后按照二分重新构建二叉树就是平衡的
func BalanceBST(root *TreeNode) *TreeNode {
	vals := make([]int, 0)
	// 中序遍历获取数据
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			vals = append(vals, node.Val)
			inOrder(node.Right)
		}
	}
	inOrder(root)
	// 构建树
	var build func(int, int) *TreeNode
	build = func(s int, e int) *TreeNode {
		if s > e {
			return nil
		}
		m := s + (e-s)>>1
		node := &TreeNode{
			Val:   vals[m],
			Left:  build(s, m-1),
			Right: build(m+1, e),
		}
		return node
	}
	return build(0, len(vals)-1)
}
