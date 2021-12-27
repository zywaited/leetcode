package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	cn := head
	num := 0
	for cn != nil {
		num++
		cn = cn.Next
	}
	ans := head
	for half := num >> 1; half > 0; half-- {
		ans = ans.Next
	}
	return ans
}
