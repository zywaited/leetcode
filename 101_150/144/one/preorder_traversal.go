package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	if root != nil {
		rs = append(rs, root.Val)
		rs = append(rs, PreorderTraversal(root.Left)...)
		rs = append(rs, PreorderTraversal(root.Right)...)
	}
	return rs
}
