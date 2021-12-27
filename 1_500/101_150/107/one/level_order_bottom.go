package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	currentLevelNodes := make([]*TreeNode, 0, 4)
	nextLevelNodes := make([]*TreeNode, 0, 4)
	currentLevelNodes = append(currentLevelNodes, root)
	for len(currentLevelNodes) > 0 {
		levelNodeValues := make([]int, 0, len(currentLevelNodes))
		for _, cn := range currentLevelNodes {
			if cn.Left != nil {
				nextLevelNodes = append(nextLevelNodes, cn.Left)
			}
			if cn.Right != nil {
				nextLevelNodes = append(nextLevelNodes, cn.Right)
			}
			levelNodeValues = append(levelNodeValues, cn.Val)
		}
		currentLevelNodes, nextLevelNodes = nextLevelNodes, currentLevelNodes[:0]
		ans = append(ans, levelNodeValues)
	}
	// 位置转换
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
