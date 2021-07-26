package one

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
		t.nums[index] = max(t.nums[index], value)
	}
}

func (t *tree) max(index int) int {
	value := 0
	for ; index > 0; index -= t.lowBit(index) {
		value = max(value, t.nums[index])
	}
	return value
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func MinOperations(target []int, arr []int) int {
	maxes := newTree(len(target) + 1)
	indexes := make(map[int]int, len(target))
	for index, num := range target {
		indexes[num] = index + 1
	}
	for _, num := range arr {
		if indexes[num] > 0 {
			maxes.update(indexes[num], maxes.max(indexes[num]-1)+1)
		}
	}
	return len(target) - maxes.max(len(target))
}
