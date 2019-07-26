package ignore_two

func CanJump(nums []int) bool {
	// 已经判断过的节点
	used := make(map[int]bool)
	queue := []int{0}
	for len(queue) != 0 {
		index := queue[0]
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
				return true
			}
			// 如果已经用过了
			if used[i] {
				continue
			}
			// 加入队列
			used[i] = true
			queue = append(queue, i)
		}
	}
	if len(nums) <= 1 {
		return true
	}
	return false
}
