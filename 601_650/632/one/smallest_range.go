package one

import (
	"container/heap"
	"math"
)

func SmallestRange(nums [][]int) []int {
	left := 0
	right := math.MinInt32
	diff := math.MaxInt32
	m := newMrh(nums)
	for i := range nums {
		if nums[i][0] > right {
			right = nums[i][0] // 更新右区间
		}
		// 左区间以堆为准
		heap.Push(m, [2]int{i, 0})
	}
	ans := []int{left, right}
	for {
		indexes := heap.Pop(m).([2]int)
		// 判断区间是否更小(区间值小或者相等时左区间小)
		if right-nums[indexes[0]][indexes[1]] < diff ||
			(right-nums[indexes[0]][indexes[1]] == diff && nums[indexes[0]][indexes[1]] < left) {
			ans[0] = nums[indexes[0]][indexes[1]]
			ans[1] = right
			diff = ans[1] - ans[0]
		}
		// 如果是当前数组是最小值，那么左区间一定是当前值，并且区间值只会越来越大(升序数组)
		if indexes[1] == len(nums[indexes[0]])-1 {
			break
		}
		// 写入当前数组下个元素
		indexes[1]++
		heap.Push(m, indexes)
		if nums[indexes[0]][indexes[1]] > right {
			right = nums[indexes[0]][indexes[1]]
		}
	}
	return ans
}

type mrh struct {
	nums [][]int
	heap [][2]int // [2]int; 0: nums-index 1: nums-index-index
}

func newMrh(nums [][]int) *mrh {
	m := &mrh{nums: nums, heap: make([][2]int, 0, len(nums))}
	heap.Init(m)
	return m
}

func (m *mrh) Len() int {
	return len(m.heap)
}

func (m *mrh) Swap(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
}

func (m *mrh) Less(i, j int) bool {
	return m.nums[m.heap[i][0]][m.heap[i][1]] < m.nums[m.heap[j][0]][m.heap[j][1]]
}

func (m *mrh) Push(indexes interface{}) {
	m.heap = append(m.heap, indexes.([2]int))
}

func (m *mrh) Pop() interface{} {
	indexes := m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]
	return indexes
}
