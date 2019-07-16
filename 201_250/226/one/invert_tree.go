package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	st := &Stack{}
	st.Push(root)
	for st.Len() > 0 {
		node := st.Pop()
		if node == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			st.Push(node.Left)
		}

		if node.Right != nil {
			st.Push(node.Right)
		}
	}

	return root
}

type Stack struct {
	nodes []*TreeNode
	l     int
}

func (st *Stack) Push(node *TreeNode) {
	st.nodes = append(st.nodes, node)
	st.l++
}

func (st *Stack) Len() int {
	return st.l
}

func (st *Stack) Pop() *TreeNode {
	if st.l == 0 {
		return nil
	}

	st.l--
	node := st.nodes[st.l]
	st.nodes = st.nodes[:st.l]
	return node
}
