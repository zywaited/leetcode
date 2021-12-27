package two

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 当快指针走到结尾的时候，慢指针就在中间节点
func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
