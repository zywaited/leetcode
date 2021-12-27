package two

import "container/heap"

// 未实现的可用的极致优化点，如果k的数量超过一定量，可以先反向去除最小的len(arr)-k个

func GetStrongest(arr []int, k int) []int {
	// 先找到中位数
	// 为了性能都用堆
	m := ((len(arr) - 1) >> 1) + 1
	// 大顶堆先获取中位数
	ns := make(numHeap, m)
	copy(ns, arr[0:m])
	heap.Init(&ns)
	for ; m < len(arr); m++ {
		if arr[m] < ns[0] {
			heap.Pop(&ns)
			heap.Push(&ns, arr[m])
		}
	}
	// 小顶堆取K个值
	sh := &strongestHeap{
		ns: make([]int, k),
		ms: ns[0], // 中位数就是堆顶
	}
	for m = 0; m < k; m++ {
		sh.ns[m] = arr[m]
	}
	heap.Init(sh)
	for ; m < len(arr); m++ {
		if sh.less(sh.ns[0], arr[m]) {
			heap.Pop(sh)
			heap.Push(sh, arr[m])
		}
	}
	return sh.ns
}

type numHeap []int

func (nh *numHeap) Len() int {
	return len(*nh)
}

func (nh *numHeap) Less(i, j int) bool {
	return (*nh)[i] > (*nh)[j]
}

func (nh *numHeap) Swap(i, j int) {
	(*nh)[i], (*nh)[j] = (*nh)[j], (*nh)[i]
}

func (nh *numHeap) Pop() interface{} {
	n := (*nh)[nh.Len()-1]
	*nh = (*nh)[:nh.Len()-1]
	return n
}

func (nh *numHeap) Push(n interface{}) {
	*nh = append(*nh, n.(int))
}

type strongestHeap struct {
	ns []int
	ms int
}

func (sh *strongestHeap) Len() int {
	return len(sh.ns)
}

func (sh *strongestHeap) Less(i, j int) bool {
	return sh.less(sh.ns[i], sh.ns[j])
}

func (sh *strongestHeap) Swap(i, j int) {
	sh.ns[i], sh.ns[j] = sh.ns[j], sh.ns[i]
}

func (sh *strongestHeap) Pop() interface{} {
	n := sh.ns[sh.Len()-1]
	sh.ns = sh.ns[:sh.Len()-1]
	return n
}

func (sh *strongestHeap) Push(n interface{}) {
	sh.ns = append(sh.ns, n.(int))
}

func (sh *strongestHeap) less(fn, sn int) bool {
	fa := abs(fn - sh.ms)
	sa := abs(sn - sh.ms)
	return (fa < sa) || (fa == sa && fn < sn)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
