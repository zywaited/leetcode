package two

// 堆
// O(N*lg(K))
func TopKFrequent(nums []int, k int) []int {
	type node struct {
		num   int
		times int
	}
	// 初始化出现频率
	mp := make(map[int]int)
	nodes := make([]node, 0)
	for _, num := range nums {
		if i, ok := mp[num]; ok {
			nodes[i].times++
			continue
		}
		nodes = append(nodes, node{num: num, times: 1})
		mp[num] = len(nodes) - 1
	}
	// 初始化数量为K的最小堆
	// 调整堆
	var heapify func(int)
	heapify = func(i int) {
		l := (i << 1) + 1
		r := (i << 1) + 2
		// 3者找到最小的
		min := i
		if l < k && nodes[l].times < nodes[min].times {
			min = l
		}
		if r < k && nodes[r].times < nodes[min].times {
			min = r
		}
		// 交换并继续往下
		if min != i {
			nodes[i], nodes[min] = nodes[min], nodes[i]
			heapify(min)
		}
	}
	for i := (k >> 1) - 1; i >= 0; i-- {
		heapify(i)
	}
	// 依次遍历与最小堆比较
	for i := k; i < len(nodes); i++ {
		if nodes[i].times > nodes[0].times {
			nodes[0] = nodes[i]
			heapify(0) // 重新调整最小堆
		}
	}
	// 获取 TOP K
	rs := make([]int, 0, k)
	for i := k - 1; i >= 0; i-- {
		rs = append(rs, nodes[i].num)
	}
	return rs
}
