package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	var (
		resp   *ListNode
		tail   *ListNode
		right  *ListNode
		remove bool
	)
	for head != nil {
		right = head.Next
		remove = false
		// 直到不同的节点
		for right != nil {
			if head.Val != right.Val {
				break
			}
			remove = true
			right = right.Next
		}
		// 去除相同的节点
		if !remove {
			if tail == nil {
				resp = head
			} else {
				tail.Next = head
			}
			tail = head
		}
		head = right
	}
	// 断链
	if tail != nil {
		tail.Next = nil
	}
	return resp
}
