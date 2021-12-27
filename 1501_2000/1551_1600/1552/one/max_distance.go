package one

import "sort"

func MaxDistance(position []int, m int) int {
	sort.Ints(position)
	left := 1                                        // 最小的磁力
	right := position[len(position)-1] - position[0] // 最大的磁力
	for left <= right {
		mid := left + (right-left)>>1
		if check(position, mid, m) {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return right
}

func check(position []int, t, m int) bool {
	n := 1              // 当前磁力t下可以放置的球数
	prev := position[0] // 前一个放置的球
	for index := 1; index < len(position); index++ {
		if position[index]-prev >= t {
			n++
			prev = position[index]
		}
	}
	return n >= m
}
