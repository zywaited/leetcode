package two

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 因为需要回溯，需要保证两个链表的长度是一致的
	max1, max2 := 0, 0
	tail1, tail2 := l1, l2
	for tail1 != nil || tail2 != nil {
		if tail1 != nil {
			max1++
			tail1 = tail1.Next
		}
		if tail2 != nil {
			max2++
			tail2 = tail2.Next
		}
	}
	// 补全l1、l2
	for max1 < max2 {
		max1++
		l1 = &ListNode{Next: l1}
	}
	for max2 < max1 {
		max2++
		l2 = &ListNode{Next: l2}
	}
	c, head := bc(l1, l2)
	if c > 0 {
		head = &ListNode{Val: c, Next: head}
	}
	return head
}

// 返回了进位和当前值
func bc(l1, l2 *ListNode) (int, *ListNode) {
	// 前面补全过，所以这里一定两个都为空或者都不为空
	if l1 == nil || l2 == nil {
		return 0, nil
	}
	sum := l1.Val + l2.Val
	c, v := bc(l1.Next, l2.Next)
	sum += c
	return sum / 10, &ListNode{Val: sum % 10, Next: v}
}
