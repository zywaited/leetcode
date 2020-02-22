package one

// 最大堆减少排序
func MinSetSize(arr []int) int {
	nums := make(map[int]int, len(arr))
	for _, num := range arr {
		nums[num]++
	}
	nodes := make([]int, 0)
	for _, num := range nums {
		// 过滤数量为1的数据
		if num > 1 {
			nodes = append(nodes, num)
		}
	}
	var heapify func(int)
	heapify = func(i int) {
		l := (i << 1) + 1
		r := l + 1
		max := i
		if l < len(nodes) && nodes[l] > nodes[max] {
			max = l
		}
		if r < len(nodes) && nodes[r] > nodes[max] {
			max = r
		}
		if max != i {
			nodes[i], nodes[max] = nodes[max], nodes[i]
			heapify(max)
		}
	}
	for i := len(nodes) - 1; i >= 0; i-- {
		heapify(i)
	}
	ans := 0
	mid := len(arr) >> 1
	for mid > 0 && len(nodes) > 0 && nodes[0] > 0 {
		ans++
		mid -= nodes[0]
		nodes[0] = 0
		heapify(0)
	}
	if mid > 0 {
		ans += mid
	}
	return ans
}
