package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		odd := head
		even := odd.Next
		node := even.Next
		next := (*ListNode)(nil)
		for node != nil {
			next = node.Next
			odd.Next, node.Next = node, odd.Next
			even.Next = next
			odd = node
			even = next

			node = nil
			if next != nil {
				node = next.Next
			}
		}
	}
	return head
}
