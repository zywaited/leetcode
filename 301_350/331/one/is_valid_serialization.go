package one

import "strings"

// 通用验证是否满足二叉树定义
func IsValidSerialization(preorder string) bool {
	// 二叉树的空节点=节点数+1
	leaveNum := 0
	nodeNum := 0
	nodes := strings.Split(preorder, ",")
	for index, node := range nodes {
		// 计算空节点和非空节点数量
		if node == "#" {
			leaveNum++
		} else {
			nodeNum++
		}
		if leaveNum > nodeNum+1 {
			return false

		}
		// 满足二叉树条件后，数组数据也应该为空
		if leaveNum == nodeNum+1 && index < len(nodes)-1 {
			return false
		}

	}
	return leaveNum == nodeNum+1
}
