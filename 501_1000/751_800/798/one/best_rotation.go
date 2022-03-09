package one

func BestRotation(nums []int) int {
	max := 0
	tr := newTree(len(nums) + 1)
	for index, num := range nums {
		if index < num {
			// 左边移动过去
			tr.update(index+1, 1)
			// 右边剩余不能少于num-index个数
			tr.update(len(nums)-(num-index)+1, -1)
			continue
		}
		// 左边可以不移动
		max++
		if num == 0 {
			tr.update(1, 1)
			continue
		}
		if num < index {
			// 左边剩余不能少于num个
			tr.update(1, 1)
			tr.update(index-num+1, -1)
		}
		// 左边全部移动
		tr.update(index+1, 1)
	}
	ans := 0
	for index := 1; index < len(nums); index++ {
		sum := tr.sum(index)
		if max < sum {
			max = sum
			ans = index
		}
	}
	return ans
}

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
