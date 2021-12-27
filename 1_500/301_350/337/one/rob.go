package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 注意点：
// 不是隔层计算而是不相连，也就是当前节点和其兄弟节点的子节点可以取
// 如果是隔层的话可以BFS+DP

// 当前节点能取的最大值就是比较取兄弟节点或兄弟节点的子节点的值
// DFS+DP
func Rob(root *TreeNode) int {
	rs := dp(root)
	return max(rs[0], rs[1])
}

// 返回值
// 0: 取子节点
// 1: 取孙节点
func dp(cn *TreeNode) [2]int {
	rs := [2]int{}
	if cn == nil {
		return rs
	}
	left := dp(cn.Left)
	right := dp(cn.Right)
	rs[0] = max(left[0], left[1]) + max(right[0], right[1]) // 子节点层最大值相加
	rs[1] = cn.Val + left[0] + right[0]                     // 隔层值相加
	return rs
}

func max(f, s int) int {
	if f < s {
		return s
	}
	return f
}
