package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	rs := &ListNode{}
	swap(rs, head)
	return rs.Next
}

// 递归，两两交换
func swap(prev, cn *ListNode) {
	if cn == nil {
		return
	}
	if cn.Next == nil {
		prev.Next = cn
		return
	}
	tmp := cn.Next
	cn.Next = tmp.Next
	tmp.Next = cn
	prev.Next = tmp
	swap(cn, cn.Next)
}
