package two

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	var (
		rs   *ListNode
		node *ListNode // 上一次交换后的尾节点
	)
	for head != nil {
		next := head.Next
		if next != nil {
			head.Next = next.Next
			if rs == nil {
				rs = next
			}
			next.Next = head
			if node != nil {
				node.Next = next
			}
			node = head
		}
		// 可能没有下一个
		if rs == nil {
			rs = head
		}
		head = head.Next
	}
	return rs
}
