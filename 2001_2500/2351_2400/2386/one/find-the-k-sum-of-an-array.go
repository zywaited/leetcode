package one

import "container/heap"

func kSum(nums []int, k int) int64 {
	sum := 0
	for i := range nums {
		if nums[i] >= 0 {
			sum += nums[i]
			continue
		}
		nums[i] = -nums[i]
	}
	sort.Ints(nums)
	q := minHeap{{nums[0], 0}}
	var d int
	var i int
	for j := 2; j <= k; j++ {
		v := heap.Pop(&q).([]int)
		d, i = v[0], v[1]
		if i == len(nums)-1 {
			continue
		}
		heap.Push(&q, []int{d + nums[i+1], i + 1})
		heap.Push(&q, []int{d + nums[i+1] - nums[i], i + 1})
	}
	return int64(sum - d)
}

type minHeap [][]int

func (heap *minHeap) Len() int {
	return len(*heap)
}

func (heap *minHeap) Less(i, j int) bool {
	return (*heap)[i][0] < (*heap)[j][0]
}

func (heap *minHeap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *minHeap) Push(value any) {
	*heap = append(*heap, value.([]int))
}

func (heap *minHeap) Pop() any {
	num := (*heap)[heap.Len()-1]
	*heap = (*heap)[:heap.Len()-1]
	return num
}
