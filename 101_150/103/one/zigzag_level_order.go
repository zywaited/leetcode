package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigzagLevelOrder(root *TreeNode) [][]int {
	rs := make([][]int, 0)
	if root != nil {
		left := true                      // 从左往右
		currentLevel := []*TreeNode{root} // 当前层
		currentLevelNum := 1              // 当前层的节点数
		nextLevel := make([]*TreeNode, 0) // 下一层
		level := 0                        // 第几层
		index := 0
		for index < currentLevelNum {
			if len(rs) <= level {
				// 初始化
				rs = append(rs, make([]int, currentLevelNum))
			}
			if left {
				rs[level][index] = currentLevel[index].Val
			} else {
				rs[level][currentLevelNum-index-1] = currentLevel[index].Val
			}
			if currentLevel[index].Left != nil {
				nextLevel = append(nextLevel, currentLevel[index].Left)
			}
			if currentLevel[index].Right != nil {
				nextLevel = append(nextLevel, currentLevel[index].Right)
			}
			// 是否当前层已经结束
			if index == currentLevelNum-1 && len(nextLevel) > 0 {
				currentLevel = nextLevel
				nextLevel = make([]*TreeNode, 0)
				currentLevelNum = len(currentLevel)
				index = 0
				left = !left
				level++
				continue
			}
			index++
		}
	}
	return rs
}
