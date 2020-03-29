package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 使用过的层级
	used := make(map[int]byte)
	type un struct {
		level int
		node  *TreeNode
	}
	queue := []*un{{level: 0, node: root}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if used[n.level] == 0 {
			ans = append(ans, n.node.Val)
			used[n.level] = 1
		}
		// 先放入右节点
		if n.node.Right != nil {
			queue = append(queue, &un{level: n.level + 1, node: n.node.Right})
		}
		if n.node.Left != nil {
			queue = append(queue, &un{level: n.level + 1, node: n.node.Left})
		}
	}
	return ans
}
