package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mp := make(map[*ListNode]byte)
	for head != nil {
		if mp[head] == 1 {
			return head
		}
		mp[head] = 1
		head = head.Next
	}
	return nil
}
