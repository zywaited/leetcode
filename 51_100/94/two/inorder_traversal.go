package two

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 栈
func InorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	curr := root
	rs := make([]int, 0)
	for curr != nil || len(stack) > 0 {
		// 如果当前节点不为空，则一直读左节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		rs = append(rs, curr.Val)
		stack = stack[:len(stack)-1]
		// 右节点重复该过程
		curr = curr.Right
	}
	return rs
}
