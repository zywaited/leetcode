package two

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	queue := []*TreeNode{root}
	prev := (*TreeNode)(nil)
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if node == nil {
			continue
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		node.Left = nil
		node.Right = nil
		if prev != nil {
			prev.Right = node
		}
		prev = node
	}
}
