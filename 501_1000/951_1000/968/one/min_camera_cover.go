package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinCameraCover(root *TreeNode) int {
	// 一个节点的状态
	// 1: 当前节点安装了监控
	// 2: 当前没有安装监控但是能被监控
	// 3: 当前节点不能被监控
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		ls, rs := dfs(node.Left), dfs(node.Right)
		if ls == 3 || rs == 3 {
			// 有一个节点没有被监控
			ans++
			return 1
		}
		if ls == 1 || rs == 1 {
			// 被子节点监控
			return 2
		}
		// 不能被监控
		return 3
	}
	if dfs(root) == 3 {
		// 监控根节点
		ans++
	}
	return ans
}
