package one

import "sort"

const mod = 1e9 + 7

func MaxPerformance(n int, speed []int, efficiency []int, k int) int {
	// 团队表现值依赖于两个指标, 这样复杂度会很高, 我们可以把效率递减排序
	// 这样就只依赖于速度和
	ses := make([][2]int, n)
	for i := 0; i < n; i++ {
		ses[i] = [2]int{speed[i], efficiency[i]}
	}
	sort.Slice(ses, func(i, j int) bool {
		return ses[i][1] > ses[j][1]
	})
	// 最小堆
	// 建立K个速度，这样就可以维持最大和
	sum := uint64(0)
	ans := uint64(0) // 可能结果很大
	nodes := make(mh, 0, k)
	for _, se := range ses {
		if len(nodes) == k {
			// 要用第i个数据，sum要减去上一个最小的
			sum -= uint64(nodes.pop())
		}
		sum += uint64(se[0])
		ans = max(ans, sum*uint64(se[1]))
		nodes.push(se[0])
	}
	return int(ans % mod)
}

func max(f, s uint64) uint64 {
	if f >= s {
		return f
	}
	return s
}

// 最小堆
type mh []int

func (ns *mh) heapify(i int, ll int) bool {
	l := i<<1 + 1
	r := l + 1
	m := i
	if l < ll && (*ns)[l] < (*ns)[m] {
		m = l
	}
	if r < ll && (*ns)[r] < (*ns)[m] {
		m = r
	}
	if m != i {
		(*ns)[i], (*ns)[m] = (*ns)[m], (*ns)[i]
		ns.heapify(m, ll)
		return true
	}
	return false
}

func (ns *mh) push(node int) {
	// 从底往上开始变化
	*ns = append(*ns, node)
	for i := (len(*ns) - 2) >> 1; i >= 0; i = (i - 1) >> 1 {
		// 不需要变化就跳过
		if !ns.heapify(i, len(*ns)) {
			break
		}
	}
}

func (ns *mh) pop() int {
	cn := (*ns)[0]
	(*ns)[0] = (*ns)[len(*ns)-1]
	*ns = (*ns)[:len(*ns)-1]
	ns.heapify(0, len(*ns))
	return cn
}
