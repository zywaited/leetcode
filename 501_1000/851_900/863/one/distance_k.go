package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := [501]*TreeNode{}
	var parseParents func(node *TreeNode)
	parseParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			parseParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			parseParents(node.Right)
		}
	}
	parseParents(root)
	ans := make([]int, 0)
	var findNodes func(node, from *TreeNode, level int)
	findNodes = func(node, from *TreeNode, level int) {
		if level == k {
			ans = append(ans, node.Val)
			return
		}
		if node.Left != nil && node.Left != from {
			findNodes(node.Left, node, level+1)
		}
		if node.Right != nil && node.Right != from {
			findNodes(node.Right, node, level+1)
		}
		if parents[node.Val] != nil && parents[node.Val] != from {
			findNodes(parents[node.Val], node, level+1)
		}
	}
	findNodes(target, nil, 0)
	return ans
}
