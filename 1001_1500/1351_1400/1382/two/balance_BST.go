package two

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 通用解法
// 实现平衡二叉树
func BalanceBST(root *TreeNode) *TreeNode {
	// 因为结构体本身没有深度和父节点，所以需要自行处理
	// 获取高度(因为val值各不相同，这里就直接用val)
	dm := make(map[int]int)
	dh := func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return dm[node.Val]
	}
	uh := func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		dm[node.Val] = max(dh(node.Left), dh(node.Right)) + 1
		return dm[node.Val]
	}
	// 左右节点差
	fh := func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return dh(node.Left) - dh(node.Right)
	}
	// 左旋
	lr := func(node *TreeNode) *TreeNode {
		rn := node.Right
		node.Right = rn.Left
		rn.Left = node
		uh(node)
		uh(rn)
		return rn
	}
	// 右旋
	rr := func(node *TreeNode) *TreeNode {
		ln := node.Left
		node.Left = ln.Right
		ln.Right = node
		uh(node)
		uh(ln)
		return ln
	}
	// 平衡
	// 因为没有parent这个属性，又不打算用map
	// 递归插入数据
	var insert func(*TreeNode, int) *TreeNode
	insert = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			// 那么就在这里创建
			node = &TreeNode{Val: val}
		} else {
			if val < node.Val {
				node.Left = insert(node.Left, val)
			} else {
				node.Right = insert(node.Right, val)
			}
			// 是否平衡
			if fh(node) < -1 || fh(node) > 1 {
				if fh(node) < -1 {
					if fh(node.Right) > 0 {
						// 先右旋
						node.Right = rr(node.Right)
					}
					// 左旋
					node = lr(node)
				} else {
					if fh(node.Left) < 0 {
						// 先左旋
						node.Left = lr(node.Left)
					}
					// 右旋
					node = rr(node)
				}
			}
		}
		// 更新当前节点高度
		uh(node)
		return node
	}
	// 写入原始数据
	ans := (*TreeNode)(nil)
	queue := &list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		ce := queue.Front()
		queue.Remove(ce)
		cn := ce.Value.(*TreeNode)
		ans = insert(ans, cn.Val)
		if cn.Left != nil {
			queue.PushBack(cn.Left)
		}
		if cn.Right != nil {
			queue.PushBack(cn.Right)
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
