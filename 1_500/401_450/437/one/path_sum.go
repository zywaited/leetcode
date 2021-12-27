package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, sum int) int {
	type node struct {
		tn   *TreeNode
		path []int
		sum  int
	}
	stack := []node{{tn: root}}
	rs := 0
	for len(stack) != 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cn.tn == nil {
			continue
		}
		path := append(cn.path, cn.tn.Val)
		cn.sum += cn.tn.Val
		// 从头判断是否有组合满足
		cs := cn.sum
		for i := 0; i < len(path); i++ {
			if cs == sum {
				rs++
			}
			cs -= path[i]
		}
		// 加入
		if cn.tn.Right != nil {
			stack = append(stack, node{tn: cn.tn.Right, path: path, sum: cn.sum})
		}
		if cn.tn.Left != nil {
			stack = append(stack, node{tn: cn.tn.Left, path: path, sum: cn.sum})
		}
	}
	return rs
}
