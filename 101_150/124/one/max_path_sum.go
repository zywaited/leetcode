package one

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	m, _ := dfs(root)
	return m
}

// 返回最大值&当前节点的最大值
func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return math.MinInt32, math.MinInt32
	}
	lm, l := dfs(node.Left)
	rm, r := dfs(node.Right)
	s := node.Val
	sm := s
	if l > 0 {
		sm += l
	}
	if r > 0 {
		sm += r
	}
	m := max(max(lm, rm), sm)
	nm := max(l, r)
	if nm >= 0 {
		s += nm
	}
	return m, s
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
