package one

type ListNode struct {
	Val  int
	Next *ListNode
}

// 先算出总和
// 注意点: 会超过限制
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		sum1 []int
		sum2 []int
		max  int // 最大位数
	)
	// 算出各自的数字
	for l1 != nil || l2 != nil {
		max++
		if l1 != nil {
			sum1 = append(sum1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			sum2 = append(sum2, l2.Val)
			l2 = l2.Next
		}
	}
	// 计算总和然后逐位取出
	var (
		head *ListNode
		c    int // 进位
		ts   int // 两数之和
	)
	// 当前位置
	index := 1
	for index <= max {
		ts = c
		if len(sum1)-index >= 0 {
			ts += sum1[len(sum1)-index]
		}
		if len(sum2)-index >= 0 {
			ts += sum2[len(sum2)-index]
		}
		head = &ListNode{Val: ts % 10, Next: head}
		c = ts / 10
		index++
	}
	// 最后一位
	if c > 0 {
		head = &ListNode{Val: c, Next: head}
	}
	return head
}
