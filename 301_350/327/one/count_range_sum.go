package one

import "sort"

func CountRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, 0)
	sum := 0
	for _, num := range nums {
		sum += num
		sums = append(sums, sum, sum-lower, sum-upper)
	}
	sort.Ints(sums)
	indexes := make(map[int]int, len(sums))
	for i := 0; i < len(sums); i++ {
		if i > 0 && sums[i] == sums[i-1] {
			continue
		}
		indexes[sums[i]] = len(indexes) + 1
	}
	ans := 0
	t := newTree(len(sums))
	sum = 0
	for _, num := range nums {
		sum += num
		left, right := indexes[sum-lower], indexes[sum-upper]
		if left > right {
			left, right = right, left
		}
		ans += t.sum(right) - t.sum(left-1)
		t.add(indexes[sum])
		if sum >= lower && sum <= upper {
			ans++
		}
	}
	return ans
}

type tree []int

func newTree(n int) *tree {
	t := make(tree, n+1)
	return &t
}

func (t *tree) len() int {
	return len(*t)
}

func (t *tree) lowBit(i int) int {
	return i & -i
}

func (t *tree) add(i int) {
	for ; i < t.len(); i += t.lowBit(i) {
		(*t)[i]++
	}
}

func (t *tree) sum(i int) int {
	n := 0
	for ; i > 0; i -= t.lowBit(i) {
		n += (*t)[i]
	}
	return n
}
