package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsEvenOddTree(root *TreeNode) bool {
	levels := []*TreeNode{root}
	level := 1
	for ; len(levels) > 0; level++ {
		ll := len(levels)
		for i := 0; i < ll; i++ {
			if (level^levels[i].Val)&1 == 1 {
				return false
			}
			if levels[i].Left != nil {
				levels = append(levels, levels[i].Left)
			}
			if levels[i].Right != nil {
				levels = append(levels, levels[i].Right)
			}
			if i == 0 {
				continue
			}
			if level&1 > 0 && levels[i].Val <= levels[i-1].Val {
				return false
			}
			if level&1 == 0 && levels[i-1].Val <= levels[i].Val {
				return false
			}
		}
		levels = levels[ll:]
	}
	return true
}
