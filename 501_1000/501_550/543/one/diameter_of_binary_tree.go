package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	// 左子树的最大深度+右子树的最大深度
	// 这里用递归方便
	ans := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		d := 0
		if node != nil {
			ld := depth(node.Left)
			rd := depth(node.Right)
			d = ld
			if rd > d {
				d = rd
			}
			d++
			if ld+rd > ans {
				ans = ld + rd
			}
		}
		return d
	}
	depth(root)
	return ans
}
