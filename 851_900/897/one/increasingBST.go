package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IncreasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root, _ = dfs(root)
	return root
}

func dfs(root *TreeNode) (l, c *TreeNode) {
	if root == nil {
		return
	}
	l = root
	c = root
	if root.Left != nil {
		ll, lc := dfs(root.Left)
		l = ll
		lc.Right = root
		root.Left = nil
	}
	if root.Right != nil {
		rl, rc := dfs(root.Right)
		root.Right = rl
		c = rc
	}
	return
}
