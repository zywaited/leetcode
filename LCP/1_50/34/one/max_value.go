package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxValue(root *TreeNode, k int) int {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		nv := make([]int, k+1)
		if node == nil {
			return nv
		}
		lv := dfs(node.Left)
		rv := dfs(node.Right)
		// 先不选
		mv := 0
		for _, v := range lv {
			if v > mv {
				mv = v
			}
		}
		nv[0] += mv
		mv = 0
		for _, v := range rv {
			if v > mv {
				mv = v
			}
		}
		nv[0] += mv
		// 选择当前
		for tk := 1; tk <= k; tk++ {
			for lk := tk - 1; lk >= 0; lk-- {
				if lv[lk]+rv[tk-lk-1] > nv[tk] {
					nv[tk] = lv[lk] + rv[tk-lk-1]
				}
			}
			nv[tk] += node.Val
		}
		return nv
	}
	mv := 0
	for _, v := range dfs(root) {
		if v > mv {
			mv = v
		}
	}
	return mv
}
