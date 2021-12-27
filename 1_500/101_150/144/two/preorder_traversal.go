package two

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			rs = append(rs, curr.Val)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:(len(stack))-1]
		curr = curr.Right
	}
	return rs
}
