package one

func CountSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	if len(nums) > 0 {
		ai := len(nums) - 1
		min, max := nums[0], nums[0]
		for _, num := range nums {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		lt := NewLineTree(min, max)
		for index := len(nums) - 1; index >= 0; index-- {
			ans[ai] = lt.Sum(min, nums[index]-1)
			ai--
			lt.Update(nums[index], nums[index], 1)
		}
	}
	return ans
}

// 最开始没有考虑到有负数
// 从树状数组转换为线段树
type (
	lineTree struct {
		root *lineTreeNode
	}

	lineTreeNode struct {
		low   int
		high  int
		sum   int
		lazy  int
		left  *lineTreeNode
		right *lineTreeNode
	}
)

func NewLineTree(low, high int) *lineTree {
	return &lineTree{root: NewLineTreeNode(low, high)}
}

func (lt *lineTree) Update(low, high, sum int) {
	lt.root.update(low, high, sum)
}

func (lt *lineTree) Sum(low, high int) int {
	return lt.root.sums(low, high)
}

func NewLineTreeNode(low, high int) *lineTreeNode {
	return &lineTreeNode{low: low, high: high}
}

func (ltn *lineTreeNode) lr() {
	if ltn.low == ltn.high {
		return
	}
	if ltn.left == nil {
		ltn.left = NewLineTreeNode(ltn.low, ltn.low+(ltn.high-ltn.low)>>1)
	}
	if ltn.right == nil {
		ltn.right = NewLineTreeNode(ltn.left.high+1, ltn.high)
	}
}

func (ltn *lineTreeNode) down() {
	ltn.lr()
	if ltn.low == ltn.high {
		ltn.lazy = 0
	}
	if ltn.lazy == 0 {
		return
	}
	// 数据结构问题
	if ltn.left == nil || ltn.right == nil {
		return
	}
	ltn.left.sum += ltn.lazy
	ltn.left.lazy += ltn.lazy
	ltn.right.sum += ltn.lazy
	ltn.right.lazy += ltn.lazy
	ltn.lazy = 0
}

func (ltn *lineTreeNode) update(low, high, sum int) {
	if ltn.high < low || ltn.low > high {
		return
	}
	if ltn.low >= low && ltn.high <= high {
		ltn.sum += sum
		ltn.lazy += sum
		return
	}
	ltn.down()
	if ltn.left.high >= low {
		ltn.left.update(low, high, sum)
	}
	if ltn.right.low <= high {
		ltn.right.update(low, high, sum)
	}
	// 重置数据
	ltn.sum = ltn.left.sum + ltn.right.sum
}

func (ltn *lineTreeNode) sums(low, high int) int {
	s := 0
	if ltn.high < low || ltn.low > high {
		return s
	}
	if ltn.low >= low && ltn.high <= high {
		return ltn.sum
	}
	ltn.down()
	if ltn.left.high >= low {
		s += ltn.left.sums(low, high)
	}
	if ltn.right.low <= high {
		s += ltn.right.sums(low, high)
	}
	return s
}
