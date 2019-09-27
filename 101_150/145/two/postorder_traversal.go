package two

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	curr := root
	rs := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 判断是否已经使用过
	// 为了不改变原始的数据
	// 不然判断玩curr.Right != nil 就应该重置 curr.Right = nil
	used := make(map[*TreeNode]byte)
	for curr != nil || len(stack) > 0 {
		// 左节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		// 这里得看下右节点是否存在
		if curr.Right == nil || used[curr] == 1 {
			// 直接添加
			rs = append(rs, curr.Val)
			stack = stack[:len(stack)-1]
			curr = nil
			continue
		}
		used[curr] = 1
		curr = curr.Right
	}
	return rs
}
