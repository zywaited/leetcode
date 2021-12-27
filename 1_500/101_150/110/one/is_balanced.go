package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deep(n *TreeNode) int {
	if n == nil {
		return 0
	}
	if dp, ok := mp[n]; ok {
		return dp
	}
	ld := deep(n.Left)
	rd := deep(n.Right)
	if ld < rd {
		ld = rd
	}
	ld++
	mp[n] = ld
	return ld
}

func balance(n *TreeNode) bool {
	diff := deep(n.Left) - deep(n.Right)
	return diff >= -1 && diff <= 1
}

// 原始定义是没有深度字段，递归每次都要重复计算
// 所以中间用了map存储一下
// 如果不需要map，把这段逻辑替换IsBalanced即可
var mp map[*TreeNode]int

func cacheIsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !balance(root) {
		return false
	}
	if cacheIsBalanced(root.Left) {
		return cacheIsBalanced(root.Right)
	}
	return false
}

func IsBalanced(root *TreeNode) bool {
	mp = make(map[*TreeNode]int)
	return cacheIsBalanced(root)
}
