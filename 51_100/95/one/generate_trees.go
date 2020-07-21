package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	return generateTree(1, n)
}

func generateTree(s, e int) []*TreeNode {
	var tns []*TreeNode
	if s <= e {
		tns = make([]*TreeNode, 0)
		for n := s; n <= e; n++ {
			if n == s && n == e {
				tns = append(tns, &TreeNode{Val: n})
				continue
			}
			lns := generateTree(s, n-1)
			rns := generateTree(n+1, e)
			if len(lns) == 0 {
				for _, rn := range rns {
					tns = append(tns, &TreeNode{Val: n, Right: rn})
				}
				continue
			}
			if len(rns) == 0 {
				for _, ln := range lns {
					tns = append(tns, &TreeNode{Val: n, Left: ln})
				}
				continue
			}
			for _, ln := range lns {
				for _, rn := range rns {
					tns = append(tns, &TreeNode{Val: n, Left: ln, Right: rn})
				}
			}
		}
	}
	return tns
}
