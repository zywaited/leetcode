package one

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if visited[node.Val] != nil {
			return visited[node.Val]
		}
		cn := &Node{Val: node.Val}
		visited[node.Val] = cn
		for _, n := range node.Neighbors {
			cn.Neighbors = append(cn.Neighbors, dfs(n))
		}
		return cn
	}
	return dfs(node)
}
