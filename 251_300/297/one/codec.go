package one

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 采用最简单逻辑，打包字符串
type Codec struct {
	placeOrder          string // 这里是不存在节点时的占位符
	separationCharacter string // 分隔符
}

func Constructor() Codec {
	return Codec{
		placeOrder:          "NULL",
		separationCharacter: ",",
	}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	// BFS
	ans := strings.Builder{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		cn := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if cn == nil {
			ans.WriteString(c.placeOrder)
			ans.WriteString(c.separationCharacter)
			continue
		}
		ans.WriteString(strconv.Itoa(cn.Val))
		ans.WriteString(c.separationCharacter)
		queue.PushBack(cn.Left)
		queue.PushBack(cn.Right)
	}
	return strings.Trim(ans.String(), c.separationCharacter)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data == c.placeOrder {
		return nil
	}
	// 这里不做合法性检测
	// 图方便，就直接用标准库
	nodes := strings.SplitN(data, c.separationCharacter, 2)
	v, _ := strconv.Atoi(nodes[0])
	ans := &TreeNode{Val: v}
	if len(nodes) == 1 || len(nodes[1]) == 0 {
		return ans
	}
	ns := nodes[1]
	queue := list.New()
	queue.PushBack(ans)
	for len(ns) > 0 {
		cn := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		nodes = strings.SplitN(ns, c.separationCharacter, 3) // 这里取三个是为了直接找到左右节点
		if nodes[0] != c.placeOrder {
			lv, _ := strconv.Atoi(nodes[0])
			cn.Left = &TreeNode{Val: lv}
			queue.PushBack(cn.Left)
		}
		if len(nodes) > 1 && nodes[1] != c.placeOrder {
			rv, _ := strconv.Atoi(nodes[1])
			cn.Right = &TreeNode{Val: rv}
			queue.PushBack(cn.Right)
		}
		ns = "" // 置空防止后续逻辑出现重复
		if len(nodes) > 2 {
			ns = nodes[2]
		}
	}
	return ans
}
