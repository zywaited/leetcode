package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxProduct(root *TreeNode) int {
	// 每个节点的左右和
	sums := make([]int, 0)
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := dfs(node.Left) + dfs(node.Right) + node.Val
		// 根节点只有于总和计算
		if node == root {
			return sum
		}
		sums = append(sums, sum)
		return sum
	}
	sum := dfs(root)
	// 计算每种情况
	ans := 0
	for _, ns := range sums {
		ans = max(ans, ns*(sum-ns))
	}
	return ans % (1e9 + 7)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
