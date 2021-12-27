package one

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedListToBST(head *ListNode) *TreeNode {
	return dfs(head, nil)
}

// 递归方便
func dfs(l, r *ListNode) *TreeNode {
	if l == r {
		return nil
	}
	m := middle(l, r)
	h := &TreeNode{Val: m.Val}
	h.Left = dfs(l, m)
	h.Right = dfs(m.Next, r)
	return h
}

// 快慢指针获取中间节点
func middle(l, r *ListNode) *ListNode {
	slow := l
	fast := l
	for fast != r && fast.Next != r {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
