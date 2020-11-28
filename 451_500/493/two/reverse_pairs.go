package two

import "sort"

// 很容易就想到了区间和
// 这里就用树状数组【线段树】

type tree struct {
	nums []int
}

func newTree(num int) *tree {
	return &tree{nums: make([]int, num)}
}

func (t *tree) lowBit(index int) int {
	return index & (-index)
}

func (t *tree) update(index, value int) {
	for ; index < len(t.nums); index += t.lowBit(index) {
		t.nums[index] += value
	}
}

func (t *tree) sum(index int) int {
	n := 0
	for ; index > 0; index -= t.lowBit(index) {
		n += t.nums[index]
	}
	return n
}

func ReversePairs(nums []int) int {
	tn := make([]int, 0, len(nums)<<1)
	for _, num := range nums {
		tn = append(tn, num, num<<1)
	}
	sort.Ints(tn)
	// 离散，找有效的数字
	ni := make(map[int]int, len(tn))
	n := 1
	for j := range tn {
		if j == 0 || tn[j-1] != tn[j] {
			ni[tn[j]] = n
			n++
		}
	}

	t := newTree(n + 1)
	cn := 0
	for i, num := range nums {
		// t.sum(ni[num<<1]) 找到小于等于当前值的数量
		// i-t.sum(ni[num<<1]) 则是大于当前值的数量
		cn += i - t.sum(ni[num<<1])
		t.update(ni[num], 1)
	}
	return cn
}
