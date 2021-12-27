package one

type unionFind struct {
	indexes []int
}

func NewUnionFind(n int) *unionFind {
	indexes := make([]int, n)
	for index := 0; index < n; index++ {
		indexes[index] = index
	}
	return &unionFind{indexes: indexes}
}

func (qf *unionFind) find(index int) int {
	for qf.indexes[index] != index {
		qf.indexes[index] = qf.indexes[qf.indexes[index]]
		index = qf.indexes[index]
	}
	return qf.indexes[index]
}

func (qf *unionFind) union(firstIndex, secondIndex int) {
	qf.indexes[qf.find(firstIndex)] = qf.indexes[qf.find(secondIndex)]
}

func FriendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	ans := make([]bool, len(requests))
	qf := NewUnionFind(n)
	for index, request := range requests {
		fp, sp := qf.find(request[0]), qf.find(request[1])
		valid := true
		for _, restriction := range restrictions {
			fv, sv := qf.find(restriction[0]), qf.find(restriction[1])
			if (fv == fp || fv == sp) && (sv == fp || sv == sp) {
				valid = false
				break
			}
		}
		if valid {
			qf.union(request[0], request[1])
			ans[index] = true
		}
	}
	return ans
}
