package one

func Jump(nums []int) int {
	// 节点
	type node struct {
		index int
		step  int
	}
	// 已经判断过的节点
	used := make(map[int]bool)
	queue := []node{{0, 0}}
	for len(queue) != 0 {
		index, step := queue[0].index, queue[0].step
		queue = queue[1:]
		// 可能最大跨度超过长度
		max := nums[index] + index
		if max > len(nums)-1 {
			max = len(nums) - 1
		}
		// 开始遍历
		for i := max; i > index; i-- {
			// 找到了
			if i == len(nums)-1 {
				return step + 1
			}
			// 如果已经用过了，上级用过代表深度肯定比现在小
			if used[i] {
				continue
			}
			// 加入队列
			used[i] = true
			queue = append(queue, node{i, step + 1})
		}
	}
	return 0
}
