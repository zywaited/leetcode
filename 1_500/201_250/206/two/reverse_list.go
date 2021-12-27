package two

type ListNode struct {
	Val  int
	Next *ListNode
}

// 循环
func ReverseList(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		next := head.Next
		head.Next = node
		node = head
		head = next
	}
	return node
}
