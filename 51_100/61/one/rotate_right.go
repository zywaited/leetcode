package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 1
	tail := head
	// 计算节点数
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	// 头节点之前的节点位置
	step := n - k%n
	if step == n {
		return head
	}
	// 形成环
	tail.Next = head
	for i := 1; i < step; i++ {
		head = head.Next
	}
	// 断开环
	newHead := head.Next
	head.Next = nil
	return newHead
}
