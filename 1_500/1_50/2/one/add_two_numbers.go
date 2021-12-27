package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// ceil
	c := false
	var n *ListNode
	var r *ListNode
	sum := 10
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if c {
			sum += 1
			c = false
		}

		if sum > 9 {
			c = true
			sum -= 10
		}

		node := &ListNode{
			Val:  sum,
			Next: nil,
		}

		if n != nil {
			n.Next = node
		} else {
			r = node
		}

		n = node
	}

	if c {
		n.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return r
}
