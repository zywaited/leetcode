package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	if root != nil {
		rs = append(rs, PostorderTraversal(root.Left)...)
		rs = append(rs, PostorderTraversal(root.Right)...)
		rs = append(rs, root.Val)
	}
	return rs
}
