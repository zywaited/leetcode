package one

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	// 当前访问的节点
	ln := root
	// 开始的节点
	pf := root
	// 临时存储当前访问的节点
	tn := ln
	for ln != nil && (ln.Left != nil || ln.Right != nil) {
		for ln != nil {
			if ln.Left == nil && ln.Right == nil {
				ln = ln.Next
				continue
			}
			tn = ln
			if tn.Left != nil {
				// 如果右节点存在则用，不存在就用下一个
				tn.Left.Next = tn.Right
				if tn.Right == nil {
					// 一直找到下一个
					for ln = ln.Next; ln != nil; ln = ln.Next {
						if ln.Left == nil && ln.Right == nil {
							continue
						}
						if ln.Left != nil {
							tn.Left.Next = ln.Left
						} else {
							tn.Left.Next = ln.Right
						}
						break
					}
					continue
				}
			}
			if tn.Right != nil {
				// 一直找到下一个
				for ln = ln.Next; ln != nil; ln = ln.Next {
					if ln.Left == nil && ln.Right == nil {
						continue
					}
					if ln.Left != nil {
						tn.Right.Next = ln.Left
					} else {
						tn.Right.Next = ln.Right
					}
					break
				}
				continue
			}
			ln = ln.Next
		}
		// 找到下一层开始的节点
		ln = pf
		for ; ln != nil; ln = ln.Next {
			if ln.Left == nil && ln.Right == nil {
				continue
			}
			if ln.Left != nil {
				if ln.Left.Left != nil || ln.Left.Right != nil {
					ln = ln.Left
					break
				}
			}
			if ln.Right != nil {
				if ln.Right.Left != nil || ln.Right.Right != nil {
					ln = ln.Right
					break
				}
			}
		}
		pf = ln
	}
	return root
}
