package one

import "sort"

type unions struct {
	indexes []int
	ens     edgeNodes
}

func newUnions(ens edgeNodes, n int) *unions {
	us := &unions{indexes: make([]int, n), ens: ens}
	for index := range us.indexes {
		us.indexes[index] = index
	}
	return us
}

func (us *unions) find(index int) int {
	for us.indexes[index] != index {
		us.indexes[index] = us.indexes[us.indexes[index]]
		index = us.indexes[index]
	}
	return index
}

func (us *unions) union(firstIndex, secondIndex int) {
	us.indexes[us.find(firstIndex)] = us.indexes[us.find(secondIndex)]
}

func (us *unions) mst(en edgeNode, ignore bool) int {
	m := 0
	ns := len(us.indexes) - 1
	if !ignore && en.n >= 0 {
		us.union(en.i, en.j)
		m = en.w
		ns--
	}
	sum := 0
	for _, edge := range us.ens {
		sum += edge.w
		if (en.n < 0 || edge.n != en.n) && us.find(edge.i) != us.find(edge.j) {
			m += edge.w
			us.union(edge.i, edge.j)
			ns--
		}
	}
	if ns > 0 {
		return sum + 1 // 比总和大就行了
	}
	return m
}

type edgeNode struct {
	i int
	j int
	w int
	n int
}

type edgeNodes []edgeNode

func newEdgeNodes(edges [][]int) edgeNodes {
	ens := make(edgeNodes, 0, len(edges))
	for index, edge := range edges {
		ens = append(ens, edgeNode{i: edge[0], j: edge[1], w: edge[2], n: index})
	}
	sort.Slice(ens, func(i, j int) bool {
		return ens[i].w < ens[j].w
	})
	return ens
}

func FindCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	ens := newEdgeNodes(edges)
	us := newUnions(ens, n)
	mst := us.mst(edgeNode{n: -1}, true)
	f := make([]int, 0)
	s := make([]int, 0)
	for _, edge := range ens {
		tus := newUnions(ens, n)
		if tus.mst(edge, true) > mst {
			f = append(f, edge.n)
			continue
		}
		tus = newUnions(ens, n)
		if tus.mst(edge, false) == mst {
			s = append(s, edge.n)
		}
	}
	return [][]int{f, s}
}
