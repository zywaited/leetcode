package one

import "sort"

// 反着推数据
// 优先队列、二叉树这种结构效率高点，但为了方便这里直接排序
func IsPossible(target []int) bool {
	// 总和
	sum := 0
	for _, num := range target {
		sum += num
	}
	sorted := false
	for {
		if !sorted {
			sort.Slice(target, func(i, j int) bool {
				return target[i] > target[j]
			})
			sorted = true
		}
		// 都是1
		if target[0] == 1 {
			return true
		}
		// 不能变化的情况
		if sum == target[0] || target[0] <= sum-target[0] {
			return false
		}
		// 重置sum(因为第一个数太大，排序后还是它，所以以倍数操作)
		multi := (sum - 1) % (sum - target[0])
		sum -= target[0]
		target[0] = multi + 1
		sum += target[0]
		// 去除1减少排序
		num := len(target) - 1
		for ; num >= 0; num-- {
			if target[num] > 1 {
				break
			}
		}
		target = target[:num+1]
		// 都是1代表刚刚好
		if len(target) == 0 {
			return true
		}
		// 这里是为了减少排序
		if target[0] == 1 {
			target = target[1:]
			continue
		}
		sorted = false
	}
}
