package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	// 找到s与t相等的节点再递归处理
	return check(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}

func check(f, s *TreeNode) bool {
	if f == nil && s == nil {
		return true
	}
	if f == nil || s == nil || f.Val != s.Val {
		return false
	}
	return check(f.Left, s.Left) && check(f.Right, s.Right)
}
