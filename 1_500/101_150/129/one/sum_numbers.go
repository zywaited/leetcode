package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			sum += s
			return
		}
		s = s*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += s
			return
		}
		if node.Left != nil {
			dfs(node.Left, s)
		}
		if node.Right != nil {
			dfs(node.Right, s)
		}
	}
	dfs(root, 0)
	return sum
}
