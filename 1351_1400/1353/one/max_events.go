package one

import "sort"

// 本来用优先队列，二叉树这种效率高一点
// 这里为了实现简单，就用排序算了
func MaxEvents(events [][]int) int {
	// 已开始时间点归类
	start := make(map[int][]int, len(events))
	queue := make([]int, 0, 40001)
	max := 0 // 最后结束的时间
	for _, event := range events {
		if event[1] > max {
			max = event[1]
		}
		start[event[0]] = append(start[event[0]], event[1])
	}
	ans := 0
	sorted := true // 为了不重复排序
	// 把当天开始的会议加到队列中排序然后取出第一个
	for i := 1; i <= max; i++ {
		if len(start[i]) > 0 {
			queue = append(queue, start[i]...)
			sorted = false
		}
		// 先排序
		if !sorted {
			sort.Slice(queue, func(i, j int) bool {
				return queue[i] < queue[j]
			})
			sorted = true
		}
		// 剔除过期的会议
		num := 0
		for num < len(queue) {
			if queue[num] < i {
				num++
				continue
			}
			break
		}
		// 没有开始的会议
		if num == len(queue) {
			queue = queue[:0]
			continue
		}
		ans++
		num++ // 取最早结束的一个会议
		// 剔除当天结束的会议
		for num < len(queue) {
			if queue[num] <= i {
				num++
				continue
			}
			break
		}
		queue = queue[num:]
	}
	return ans
}
