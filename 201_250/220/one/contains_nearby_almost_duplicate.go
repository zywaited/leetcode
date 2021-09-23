package one

import "sort"

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	allNums := make([]int, 0, len(nums)<<1)
	for _, num := range nums {
		allNums = append(allNums, num, num-t, num+t)
	}
	sort.Ints(allNums)
	numIndexes := make(map[int]int, len(allNums))
	for index := range allNums {
		if index == 0 || allNums[index] != allNums[index-1] {
			numIndexes[allNums[index]] = len(numIndexes) + 1
		}
	}
	tr := newTree(len(allNums))
	for index, num := range nums {
		firstIndex := numIndexes[num-t]
		secondIndex := numIndexes[num+t]
		if secondIndex < firstIndex {
			firstIndex += secondIndex
			secondIndex = firstIndex - secondIndex
			firstIndex = firstIndex - secondIndex
		}
		if tr.sum(secondIndex)-tr.sum(firstIndex-1) > 0 {
			return true
		}
		tr.update(numIndexes[num], 1)
		if index >= k {
			tr.update(numIndexes[nums[index-k]], -1)
		}
	}
	return false
}

type tree []int

func newTree(length int) *tree {
	t := make(tree, length+1)
	return &t
}

func (t *tree) len() int {
	return len(*t)
}

func (t *tree) lowBit(index int) int {
	return index & (-index)
}

func (t *tree) update(index int, value int) {
	for ; index < t.len(); index += t.lowBit(index) {
		(*t)[index] += value
	}
}

func (t *tree) sum(index int) int {
	sum := 0
	for ; index > 0; index -= t.lowBit(index) {
		sum += (*t)[index]
	}
	return sum
}
