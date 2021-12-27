package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	rs := InorderTraversal(root.Left)
	rs = append(rs, root.Val)
	rs = append(rs, InorderTraversal(root.Right)...)
	return rs
}
