package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CountNodes(node *TreeNode) int {
	num := 0
	if node != nil {
		ld := depth(node.Left)
		rd := 0
		for node != nil {
			rd = depth(node.Right)
			if ld == rd {
				// 左子树满二叉树
				num += 1 << ld
				ld = rd - 1
				node = node.Right
				continue
			}
			// 右边
			num += 1 << rd
			ld--
			node = node.Left
		}
	}
	return num
}

func depth(node *TreeNode) int {
	d := 0
	for ; node != nil; node = node.Left {
		d++
	}
	return d
}
