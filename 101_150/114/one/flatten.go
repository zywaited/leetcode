package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	order(root)
}

// 返回最上层节点&最底层的节点
func order(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}
	nodes := make([]*TreeNode, 0, 2)
	lns := order(node.Left)
	rns := order(node.Right)
	// 先重置数组
	cn := node
	node.Left = nil
	node.Right = nil
	if len(lns) > 0 {
		cn.Right = lns[0]
		cn = lns[1]
	}
	if len(rns) > 0 {
		cn.Right = rns[0]
		cn = rns[1]
	}
	nodes = append(nodes, node)
	nodes = append(nodes, cn)
	return nodes
}
