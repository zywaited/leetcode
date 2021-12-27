package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxLevelSum(root *TreeNode) int {
	type nd struct {
		n *TreeNode
		l int // level
	}
	// 默认第一层最大
	m := root.Val
	ml := 1
	cl := 1
	cm := 0
	qe := []nd{{n: root, l: 1}}
	for len(qe) > 0 {
		tn := qe[0]
		if tn.n == nil {
			break
		}
		qe = qe[1:]
		if tn.l == cl {
			cm += tn.n.Val
		} else {
			if cm > m {
				m = cm
				ml = cl
			}
			cl = tn.l
			cm = tn.n.Val
		}
		if tn.n.Left != nil {
			qe = append(qe, nd{n: tn.n.Left, l: tn.l + 1})
		}
		if tn.n.Right != nil {
			qe = append(qe, nd{n: tn.n.Right, l: tn.l + 1})
		}
	}
	return ml
}
