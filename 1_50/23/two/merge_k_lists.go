package two

type ListNode struct {
	Val  int
	Next *ListNode
}

// 比较
func MergeKLists(lists []*ListNode) *ListNode {
	var (
		tail  *ListNode
		rs    *ListNode
		small *ListNode
		index int
	)
	for {
		// 每一轮取出最小的
		small = nil
		for i, head := range lists {
			if head != nil && (small == nil || head.Val < small.Val) {
				small = head
				index = i
			}
		}
		// 已经结尾
		if small == nil {
			break
		}
		if rs == nil {
			rs = small
		}
		if tail != nil {
			tail.Next = small
		}
		tail = small
		// 去掉第一个
		lists[index] = small.Next
	}
	return rs
}
