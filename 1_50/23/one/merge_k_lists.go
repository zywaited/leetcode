package one

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 把所有的数据取出来然后排序
func MergeKLists(lists []*ListNode) *ListNode {
	var (
		tail  *ListNode // 为了重复利用内存
		nodes []int
		head  *ListNode
	)
	for _, list := range lists {
		for list != nil {
			if tail != nil {
				tail.Next = list
			}
			tail = list
			if head == nil {
				head = tail
			}
			nodes = append(nodes, list.Val)
			list = list.Next
		}
	}
	sort.Ints(nodes)
	tail = head
	for _, node := range nodes {
		tail.Val = node
		tail = tail.Next
	}
	return head
}
