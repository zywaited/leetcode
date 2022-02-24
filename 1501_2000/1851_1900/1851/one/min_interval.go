package one

import (
	"container/heap"
	"sort"
)

func MinInterval(intervals [][]int, queries []int) []int {
	queryNodes := make([]queryNode, len(queries))
	for index, query := range queries {
		queryNodes[index] = queryNode{index: index, query: query}
	}
	sort.Slice(queryNodes, func(i, j int) bool {
		return queryNodes[i].query < queryNodes[j].query
	})
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] == intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([]int, len(queries))
	queue := prioryQueue{nodes: make([]intervalNode, 0, len(intervals))}
	intervalIndex := 0
	for _, query := range queryNodes {
		for ; intervalIndex < len(intervals); intervalIndex++ {
			if query.query < intervals[intervalIndex][0] {
				break
			}
			if intervals[intervalIndex][1] < query.query {
				continue
			}
			heap.Push(&queue, intervalNode{index: intervalIndex, length: intervals[intervalIndex][1] - intervals[intervalIndex][0] + 1})
		}
		for queue.Len() > 0 {
			if query.query < intervals[queue.nodes[0].index][0] || intervals[queue.nodes[0].index][1] < query.query {
				heap.Pop(&queue)
				continue
			}
			break
		}
		if queue.Len() == 0 {
			ans[query.index] = -1
			continue
		}
		ans[query.index] = queue.nodes[0].length
	}
	return ans
}

type queryNode struct {
	index int
	query int
}

type intervalNode struct {
	index  int
	length int
}

type prioryQueue struct {
	nodes []intervalNode
}

func (p *prioryQueue) Len() int {
	return len(p.nodes)
}

func (p *prioryQueue) Less(i, j int) bool {
	return p.nodes[i].length < p.nodes[j].length
}

func (p *prioryQueue) Swap(i, j int) {
	p.nodes[i], p.nodes[j] = p.nodes[j], p.nodes[i]
}

func (p *prioryQueue) Push(node interface{}) {
	p.nodes = append(p.nodes, node.(intervalNode))
}

func (p *prioryQueue) Pop() interface{} {
	node := p.nodes[p.Len()-1]
	p.nodes = p.nodes[:p.Len()-1]
	return node
}
