package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	ans := &ListNode{Next: head}
	m := ans
	for ; m.Next != nil && m.Next.Val < x; m = m.Next {
	}
	if m.Next == nil {
		return head
	}
	// found
	l := m.Next
	for l != nil {
		for ; l.Next != nil && l.Next.Val >= x; l = l.Next {
		}
		if l.Next == nil {
			break
		}
		m.Next, l.Next, l.Next.Next = l.Next, l.Next.Next, m.Next
		m = m.Next
	}
	return ans.Next
}
