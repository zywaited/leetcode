package one

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	if root != nil {
		q := list.New()
		q.PushBack(root.Left)
		q.PushBack(root.Right)
		for q.Len() > 0 {
			f := q.Front()
			q.Remove(f)
			s := q.Front()
			q.Remove(s)
			fn := f.Value.(*TreeNode)
			sn := s.Value.(*TreeNode)
			if (fn == nil && sn != nil) || (fn != nil && sn == nil) {
				return false
			}
			if fn == nil || sn == nil {
				continue
			}
			if fn.Val != sn.Val {
				return false
			}
			q.PushBack(fn.Left)
			q.PushBack(sn.Right)
			q.PushBack(fn.Right)
			q.PushBack(sn.Left)
		}
	}
	return true
}
