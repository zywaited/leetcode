package one

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	nm := make(map[*Node]*Node)
	for ch := head; ch != nil; ch = ch.Next {
		nm[ch] = &Node{Val: ch.Val}
	}
	tmp := &Node{}
	cn := tmp
	for ch := head; ch != nil; ch = ch.Next {
		cn.Next = nm[ch]
		cn.Next.Random = nm[ch.Random]
		cn = cn.Next
	}
	return tmp.Next
}
