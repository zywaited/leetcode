package one

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}
		return dfs(node.Left, low, node.Val) && dfs(node.Right, node.Val, high)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
