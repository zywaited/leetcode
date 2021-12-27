package two

import "strings"

// 根据前序逻辑验证
func IsValidSerialization(preorder string) bool {
	if len(preorder) > 0 {
		// 用栈解决
		stack := make([]bool, 0, len(preorder))
		nodes := strings.Split(preorder, ",")
		for index := 0; index < len(nodes); index++ {
			if nodes[index] == "#" {
				for {
					// 已经到顶点
					if len(stack) == 0 {
						return index == len(nodes)-1
					}
					if !stack[len(stack)-1] {
						// 没结束
						break
					}
					// 去除已经找到左右节点的
					stack = stack[:len(stack)-1]
				}
				// 左边已经遍历完成
				stack[len(stack)-1] = true
				continue
			}
			stack = append(stack, false)
		}
		return len(stack) == 0
	}
	return false
}
