package one

func MinJumps(arr []int) int {
	// bfs节点
	type node struct {
		index int
		step  int
	}
	// 待处理的队列
	queue := make([]*node, 0, len(arr))
	queue = append(queue, &node{})
	// 使用过的节点
	used := make(map[int]byte, len(arr))
	used[0] = 1
	// 相同节点
	eq := make(map[int][]int, len(arr))
	for index, num := range arr {
		eq[num] = append(eq[num], index)
	}
	for len(queue) > 0 {
		cn := queue[0]
		queue = queue[1:]
		if cn.index == len(arr)-1 {
			return cn.step
		}
		// 可以到的节点
		// i+1
		if cn.index+1 < len(arr) && used[cn.index+1] == 0 {
			if cn.index+1 == len(arr)-1 {
				return cn.step + 1
			}
			queue = append(queue, &node{index: cn.index + 1, step: cn.step + 1})
			used[cn.index+1] = 1
		}
		// i-1
		if cn.index-1 >= 0 && used[cn.index-1] == 0 {
			queue = append(queue, &node{index: cn.index - 1, step: cn.step + 1})
			used[cn.index-1] = 1
		}
		// arr[i] == arr[j]
		for _, index := range eq[arr[cn.index]] {
			if index == cn.index {
				continue
			}
			if index == len(arr)-1 {
				return cn.step + 1
			}
			if used[index] == 0 {
				queue = append(queue, &node{index: index, step: cn.step + 1})
				used[index] = 1
			}
		}
		eq[arr[cn.index]] = nil // 清空数据减少上诉for循环
	}
	return 0
}
