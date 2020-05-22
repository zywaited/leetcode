package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 中序遍历每个数字索引
	im := make(map[int]int, len(inorder))
	for i, n := range inorder {
		im[n] = i
	}
	root := &TreeNode{Val: preorder[0]}
	stack := make([]*TreeNode, 0, len(preorder))
	stack = append(stack, root)
	for i := 1; i < len(preorder); i++ {
		si := len(stack) - 1
		if im[preorder[i]] < im[stack[si].Val] {
			// 左边
			stack[si].Left = &TreeNode{Val: preorder[i]}
			stack = append(stack, stack[si].Left)
			continue
		}
		// 右边
		// 找到最接近的索引节点(保留根节点)
		for si-1 >= 0 && im[stack[si-1].Val] < im[preorder[i]] {
			si--
		}
		// 存在右节点后需要去除当前节点，所以cn临时保存用于赋值
		cn := stack[si]
		stack = stack[:si]
		cn.Right = &TreeNode{Val: preorder[i]}
		stack = append(stack, cn.Right)
	}
	return root
}
