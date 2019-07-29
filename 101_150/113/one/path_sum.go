package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, sum int) [][]int {
	// 这里就不用递归方式了，虽然也可以
	type node struct {
		tn   *TreeNode
		sum  int
		path []int
	}
	stack := []node{{tn: root, sum: sum}}
	var rs [][]int
	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cn.tn == nil {
			continue
		}
		// 找到了
		left, right := cn.tn.Left != nil, cn.tn.Right != nil
		if cn.tn.Val == cn.sum && !left && !right {
			rs = append(rs, append(append([]int{}, cn.path...), cn.tn.Val))
			continue
		}
		// 判断有无都可
		if right {
			stack = append(stack, node{tn: cn.tn.Right, sum: cn.sum - cn.tn.Val, path: append(cn.path, cn.tn.Val)})
		}
		if left {
			stack = append(stack, node{tn: cn.tn.Left, sum: cn.sum - cn.tn.Val, path: append(cn.path, cn.tn.Val)})
		}
	}
	return rs
}
