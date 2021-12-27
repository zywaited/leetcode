package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BtreeGameWinningMove(root *TreeNode, n int, x int) bool {
	// 二号玩家有3个取数方案
	// 1: x的左节点数
	// 2: x的右节点数
	// 3: x的父节点的另一个子节点

	// 这里找x节点和数量计算分开，不想放一起
	// 选则最大值即可
	target := (*TreeNode)(nil) // 找到x节点
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		// 已经找到了
		if tn == nil || target != nil {
			return
		}
		if tn.Val == x {
			target = tn
			return
		}
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root)
	// target不可能为空就不判断了
	var num func(*TreeNode) int
	num = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		return num(tn.Left) + num(tn.Right) + 1
	}
	l := num(target.Left)
	r := num(target.Right)
	max := n - l - r - 1
	if max < l {
		max = l
	}
	if max < r {
		max = r
	}
	return max > n-max
}
