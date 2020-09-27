package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	// 找到中序和后序遍历的分割点
	splitIndex := len(postorder)
	for index := range inorder {
		if inorder[index] == root.Val {
			splitIndex = index
			break
		}
	}
	root.Left = BuildTree(inorder[:splitIndex], postorder[:splitIndex])
	root.Right = BuildTree(inorder[splitIndex+1:], postorder[splitIndex:len(postorder)-1])
	return root
}
