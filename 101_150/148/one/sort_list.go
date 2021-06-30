package one

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	nl := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		nl++
	}
	tmpHead := &ListNode{Next: head}
	for cl := 1; cl <= nl; cl <<= 1 {
		cn := tmpHead.Next
		tn := tmpHead
		for cn != nil {
			// 两边每次取cl个节点进行合并
			ltn := cn
			rtn := cn.Next
			for tl := 1; tl < cl; tl++ {
				if ltn == nil {
					break
				}
				ltn = ltn.Next
				if rtn != nil && rtn.Next != nil {
					rtn = rtn.Next.Next
				}
			}
			if ltn == nil || ltn.Next == nil {
				tn.Next = cn
				cn = nil
				continue
			}
			tcn := rtn
			if rtn != nil {
				tcn = rtn.Next
				rtn.Next = nil
			}
			rn := ltn.Next
			ltn.Next = nil
			tn.Next, tn = merge(cn, rn)
			cn = tcn
		}
	}
	return tmpHead.Next
}

func merge(f, s *ListNode) (*ListNode, *ListNode) {
	head := &ListNode{}
	prev := head
	for f != nil && s != nil {
		if f.Val <= s.Val {
			prev.Next = f
			f = f.Next
		} else {
			prev.Next = s
			s = s.Next
		}
		prev = prev.Next
	}
	if f == nil {
		f = s
	}
	for f != nil {
		prev.Next = f
		prev = prev.Next
		f = f.Next
	}
	return head.Next, prev
}
