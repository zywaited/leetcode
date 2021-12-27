package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxAncestorDiff(root *TreeNode) int {
	ans, _, _ := ancestorDiff(root)
	return ans
}

// 返回最大差值和最大|小节点值
func ancestorDiff(node *TreeNode) (int, int, int) {
	diff := 0
	maxV := node.Val
	minV := node.Val
	if node.Left != nil {
		ld, lMaxV, lMinV := ancestorDiff(node.Left)
		diff = max(diff, ld)
		diff = max(diff, abs(lMaxV-node.Val))
		diff = max(diff, abs(node.Val-lMinV))
		maxV = max(maxV, lMaxV)
		minV = min(minV, lMinV)
	}
	if node.Right != nil {
		rd, rMaxV, rMinV := ancestorDiff(node.Right)
		diff = max(diff, rd)
		diff = max(diff, abs(rMaxV-node.Val))
		diff = max(diff, abs(node.Val-rMinV))
		maxV = max(maxV, rMaxV)
		minV = min(minV, rMinV)
	}
	return diff, maxV, minV
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
