package one

import "math"

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

func LargestComponentSize(nums []int) int {
	maxNum := 0
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	qf := NewUnionFind(maxNum + 1)
	for _, num := range nums {
		for arg := int(math.Sqrt(float64(num))); arg > 1; arg-- {
			if num%arg == 0 {
				qf.union(num, arg)
				qf.union(num, num/arg)
			}
		}
	}
	ans := 0
	qfs := make([]int, maxNum+1)
	for _, num := range nums {
		index := qf.find(num)
		qfs[index]++
		ans = max(ans, qfs[index])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
