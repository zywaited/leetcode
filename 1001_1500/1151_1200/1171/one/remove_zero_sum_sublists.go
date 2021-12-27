package one

type ListNode struct {
	Val  int
	Next *ListNode
}

// 包含自身的前缀和
func RemoveZeroSumSublists(head *ListNode) *ListNode {
	prefix := make(map[int]*ListNode) // 每个和对应的节点,方便删除
	th := &ListNode{Next: head}
	prefix[0] = th // 第一个前缀和为0
	sum := 0
	tmp := (*ListNode)(nil)
	tmpSum := sum
	for head != nil {
		sum += head.Val
		// 判断前缀和是否已经存在
		if prefix[sum] != nil {
			// 去除中间的前缀和以及节点
			tmpSum = sum
			tmp = prefix[sum].Next
			prefix[sum].Next = head.Next
			for tmp != head {
				tmpSum += tmp.Val
				prefix[tmpSum] = nil
				tmp = tmp.Next
			}
		} else {
			prefix[sum] = head
		}
		head = head.Next
	}
	return th.Next
}
