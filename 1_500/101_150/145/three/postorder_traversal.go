package three

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	if root != nil {
		nodes := list.List{}
		stack := []*TreeNode{root}
		cn := (*TreeNode)(nil)
		for len(stack) > 0 {
			cn = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nodes.PushFront(cn.Val) // 每次都放到前面，因为是访问根、右、左，这里用双向链表是为了不频繁移动数组
			if cn.Left != nil {
				stack = append(stack, cn.Left)
			}
			if cn.Right != nil {
				stack = append(stack, cn.Right)
			}
		}
		ce := nodes.Front()
		for ce != nil {
			rs = append(rs, ce.Value.(int))
			ce = ce.Next()
		}
	}
	return rs
}
