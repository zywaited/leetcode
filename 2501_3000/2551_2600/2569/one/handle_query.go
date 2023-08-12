package one

func HandleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	// num2Sum := num2PrevSum + num1Sum * p
	num2Sum := int64(0)
	for _, num := range nums2 {
		num2Sum += int64(num)
	}
	var sums []int64
	tree := newSegmentTree(nums1)
	for _, query := range queries {
		switch query[0] {
		case 1:
			tree.revert(query[1]+1, query[2]+1)
		case 2:
			num2Sum += tree.sum * int64(query[1])
		case 3:
			sums = append(sums, num2Sum)
		}
	}
	return sums
}

// 线段树

func newSegmentTree(nums []int) *segmentTree {
	prevSums := make([]int64, len(nums)+1)
	for index, num := range nums {
		prevSums[index+1] = prevSums[index] + int64(num)
	}
	return &segmentTree{
		low:      1,
		high:     len(nums),
		sum:      prevSums[len(nums)],
		prevSums: prevSums,
	}
}

type segmentTree segmentNode

func (tree *segmentTree) revert(low, high int) {
	(*segmentNode)(tree).revert(low, high)
}

type segmentNode struct {
	low      int
	high     int
	times    int
	sum      int64
	prevSums []int64
	left     *segmentNode
	right    *segmentNode
}

func (node *segmentNode) revert(low, high int) {
	if high < node.low || node.high < low {
		return
	}
	if low <= node.low && node.high <= high {
		node.sum = int64(node.high-node.low+1) - node.sum
		node.times++
		return
	}
	node.lazyInitChildren()
	node.left.revert(low, high)
	node.right.revert(low, high)
	node.sum = node.left.sum + node.right.sum
	node.times = 0
}

func (node *segmentNode) lazyInitChildren() {
	if node.left != nil && node.right != nil {
		node.left.revertSum(node.times)
		node.right.revertSum(node.times)
		return
	}
	mid := node.low + (node.high-node.low)/2
	node.left = &segmentNode{
		low:      node.low,
		high:     mid,
		times:    node.times,
		prevSums: node.prevSums,
	}
	node.left.lazyInitSum()
	node.right = &segmentNode{
		low:      mid + 1,
		high:     node.high,
		times:    node.times,
		prevSums: node.prevSums,
	}
	node.right.lazyInitSum()
}

func (node *segmentNode) lazyInitSum() {
	sum := node.prevSums[node.high] - node.prevSums[node.low-1]
	if node.times%2 == 1 {
		sum = int64(node.high-node.low+1) - sum
	}
	node.sum = sum
}

func (node *segmentNode) revertSum(lazyTimes int) {
	node.times += lazyTimes
	if lazyTimes%2 == 1 {
		node.sum = int64(node.high-node.low+1) - node.sum
	}
}
