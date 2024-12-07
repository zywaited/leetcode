package one

import "github.com/emirpasic/gods/trees/redblacktree"

func oddEvenJumps(nums []int) int {
	tree := redblacktree.NewWithIntComparator()
	answer := 1
	// 记录每个索引能否到达尾部
	type markNode struct {
		odd  bool
		even bool
	}
	marks := make([]markNode, len(nums))
	initIndex := len(nums) - 1
	marks[initIndex].odd = true
	marks[initIndex].even = true
	tree.Put(nums[initIndex], initIndex)
	for i := len(nums) - 2; i >= 0; i-- {
		treeNode, ok := tree.Ceiling(nums[i])
		if ok {
			marks[i].odd = marks[treeNode.Value.(int)].even
		}
		treeNode, ok = tree.Floor(nums[i])
		if ok {
			marks[i].even = marks[treeNode.Value.(int)].odd
		}
		tree.Put(nums[i], i)
		if marks[i].odd {
			answer++
		}
	}
	return answer
}
