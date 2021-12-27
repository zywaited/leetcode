package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	// 这种情况不需要反转
	if k <= 1 {
		return head
	}
	// 减少内存使用先循环一次
	// 先找出K的倍数
	// 这样在最后就不用做任何操作了
	multi := 0
	th := head
	for th != nil {
		multi++
		th = th.Next
	}
	multi /= k
	return reverse(head, k, multi)
}

// 反转K个节点
func reverse(head *ListNode, k, multi int) *ListNode {
	// 剩余节点不用转
	if multi == 0 {
		return head
	}
	num := 1
	node := head
	var (
		prev *ListNode
		tail *ListNode
	)
	for {
		// 反转
		tail = node.Next
		node.Next = prev
		if num < k {
			prev = node
			node = tail
			num++
			continue
		}
		head.Next = reverse(tail, k, multi-1)
		return node
	}
}
