package one

import "container/heap"

func kClosest(points [][]int, K int) [][]int {
	ns := make(nodes, 0, K)
	for i := 0; i < K; i++ {
		ns = append(ns, node{
			point: points[i],
			v:     points[i][0]*points[i][0] + points[i][1]*points[i][1],
		})
	}
	heap.Init(&ns)
	for i := K; i < len(points); i++ {
		v := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		if ns[0].v > v {
			n := heap.Pop(&ns).(node)
			n.point = points[i]
			n.v = v
			heap.Push(&ns, n)
		}
	}
	ans := make([][]int, 0, ns.Len())
	for _, n := range ns {
		ans = append(ans, n.point)
	}
	return ans
}

type node struct {
	point []int
	v     int
}

type nodes []node

func (ns *nodes) Len() int {
	return len(*ns)
}

func (ns *nodes) Less(i, j int) bool {
	return (*ns)[i].v > (*ns)[j].v
}

func (ns *nodes) Swap(i, j int) {
	(*ns)[i], (*ns)[j] = (*ns)[j], (*ns)[i]
}

func (ns *nodes) Push(v interface{}) {
	*ns = append(*ns, v.(node))
}

func (ns *nodes) Pop() interface{} {
	v := (*ns)[ns.Len()-1]
	*ns = (*ns)[:ns.Len()-1]
	return v
}
