package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConvertBST(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, prev int) int
	dfs = func(node *TreeNode, prev int) int {
		if node == nil {
			return 0
		}
		sum := dfs(node.Right, prev) + node.Val
		node.Val = sum + prev
		if node.Left != nil {
			ls := dfs(node.Left, node.Val)
			sum += ls
		}
		return sum
	}
	dfs(root, 0)
	return root
}
