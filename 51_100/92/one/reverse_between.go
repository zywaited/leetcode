package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, f int, s int) *ListNode {
	if head != nil && f < s {
		// 找到m节点
		i := 1
		m := head
		mp := (*ListNode)(nil) // m prev
		for ; i < f; i++ {
			mp = m
			m = m.Next
		}
		tn := m
		p := (*ListNode)(nil)
		for ; i < s; i++ {
			n := tn.Next
			tn.Next = p
			p = tn
			tn = n
		}
		m.Next = tn.Next
		tn.Next = p
		if mp != nil {
			mp.Next = tn
		} else {
			head = tn
		}
	}
	return head
}
