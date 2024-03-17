package one

import "container/heap"

func maxMoves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ts := make([][]int, m)
	mh := minHeap{grid: grid}
	for i := range grid {
		ts[i] = make([]int, n)
		mh.nodes = append(mh.nodes, minHeapNode{i: i})
	}
	heap.Init(&mh)
	mt := 0
	for mh.Len() > 0 {
		node := heap.Pop(&mh).(minHeapNode)
		i, j := node.i, node.j
		if mt < ts[i][j] {
			mt = ts[i][j]
		}
		for _, direct := range [][]int{{-1, 1}, {0, 1}, {1, 1}} {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || m <= ni || nj < 0 || n <= nj || grid[ni][nj] <= grid[i][j] {
				continue
			}
			pt := ts[ni][nj]
			if ts[ni][nj] < ts[i][j]+1 {
				ts[ni][nj] = ts[i][j] + 1
			}
			if nj == 0 || pt > 0 {
				continue
			}
			heap.Push(&mh, minHeapNode{i: ni, j: nj})
		}
	}
	return mt
}

type minHeapNode struct {
	i int
	j int
}

type minHeap struct {
	grid  [][]int
	nodes []minHeapNode
}

func (mh *minHeap) Len() int {
	return len(mh.nodes)
}

func (mh *minHeap) Less(i, j int) bool {
	f, s := mh.nodes[i], mh.nodes[j]
	return mh.grid[f.i][f.j] < mh.grid[s.i][s.j]
}

func (mh *minHeap) Swap(i, j int) {
	mh.nodes[i], mh.nodes[j] = mh.nodes[j], mh.nodes[i]
}

func (mh *minHeap) Push(value interface{}) {
	mh.nodes = append(mh.nodes, value.(minHeapNode))
}

func (mh *minHeap) Pop() interface{} {
	node := mh.nodes[mh.Len()-1]
	mh.nodes = mh.nodes[:mh.Len()-1]
	return node
}
