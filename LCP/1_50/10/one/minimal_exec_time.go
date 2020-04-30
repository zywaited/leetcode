package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左子树, 总和: s1, 并行: d1, 剩余时间: s1-d1*2
// 右子树, 总和: s2, 并行: d2, 剩余时间: s2-d2*2
// 最后剩余时间继续并行, 那么所需就是两个的较大值
// 即当前节点所需时间: val + d1 + d2 + max(s1-d1*2, s2-d2*2)
// 转换公式: val + max(s1+d2-d1, s2+d1-d2) = val + max(s1-(d1-d2), s2+d1-d2)
// 数学：两条线相交时, max这时就是最小的
// 即s1-(d1-d2) = s2+d1-d2, d1-d2 = (s1-s2)/2, time = s1-(s1-s2)/2 = (s1+s2)/2
// 还有两种可能是无法相交(s2+d1-d2取值范围不能小于d2, 因为d2已经是最优解, 同理s1-(d1-d2)不能小于d1), 所以实际取的最大值3种情况的最大值
func MinimalExecTime(root *TreeNode) float64 {
	var dfs func(*TreeNode) (float64, float64)
	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0, 0
		}
		s1, d1 := dfs(node.Left)
		s2, d2 := dfs(node.Right)
		// 总时间
		s := float64(node.Val) + s1 + s2
		// 总体并行时间
		d := (s1 + s2) / float64(2)
		if d1 > d {
			d = d1
		}
		if d2 > d {
			d = d2
		}
		return s, d + float64(node.Val)
	}
	_, ans := dfs(root)
	return ans
}
