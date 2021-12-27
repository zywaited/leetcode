package one

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// a.next = b
// a.left.next = a.right
// a.right.next = a.next.left
func Connect(root *Node) *Node {
	// 每一层正在处理的节点
	ln := root
	// 每一层开始的节点
	lf := (*Node)(nil)
	if ln != nil {
		lf = ln.Left
	}
	// 完美二叉树(下一层是否存在)
	for lf != nil {
		for ln != nil {
			ln.Left.Next = ln.Right
			if ln.Next != nil {
				ln.Right.Next = ln.Next.Left
			}
			ln = ln.Next
		}
		ln = lf
		lf = ln.Left
	}
	return root
}
