package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	over := root.Left == nil && root.Right == nil
	if root.Val != sum && over {
		return false
	}
	if root.Val == sum && over {
		return true
	}
	if HasPathSum(root.Left, sum-root.Val) {
		return true
	}
	return HasPathSum(root.Right, sum-root.Val)
}
