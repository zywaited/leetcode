package one

import (
	"container/heap"
	"container/list"
)

const max = 10

type node struct {
	tweetId int
	weight  int
}

type Twitter struct {
	// 每个用户关注列表
	fl map[int]*list.List
	fe map[int]map[int]*list.Element // 这个是为了取消关注O(1)
	// 每个用户发的信息(这里不涉及删除，就用数组)
	tl     map[int][]*node
	weight int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		fl: make(map[int]*list.List),
		fe: make(map[int]map[int]*list.Element),
		tl: make(map[int][]*node),
	}
}

func (tr *Twitter) GetNFeed(userId, num int) []*node {
	if num >= len(tr.tl[userId]) {
		return tr.tl[userId]
	}
	return tr.tl[userId][len(tr.tl[userId])-num:]
}

/** Compose a new tweet. */
func (tr *Twitter) PostTweet(userId int, tweetId int) {
	tr.weight++
	tr.tl[userId] = append(tr.tl[userId], &node{tweetId: tweetId, weight: tr.weight})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (tr *Twitter) GetNewsFeed(userId int) []int {
	nodes := make(intHp, 0, max)
	nodes = append(nodes, tr.GetNFeed(userId, max)...)
	heap.Init(&nodes)
	// 每个关注者取对应的数据
	if tr.fl[userId] != nil && tr.fl[userId].Len() > 0 {
		ce := tr.fl[userId].Back()
		for ; ce != nil; ce = ce.Prev() {
			// 关注了自己
			if ce.Value.(int) == userId {
				continue
			}
			for _, node := range tr.GetNFeed(ce.Value.(int), max) {
				if nodes.Len() == max && nodes.Top().weight >= node.weight {
					continue
				}
				if nodes.Len() == max {
					heap.Pop(&nodes)
				}
				heap.Push(&nodes, node)
			}
		}
	}
	ls := make([]int, nodes.Len())
	for nodes.Len() > 0 {
		ls[nodes.Len()-1] = heap.Pop(&nodes).(*node).tweetId
	}
	return ls
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (tr *Twitter) Follow(followerId int, followeeId int) {
	if tr.fe[followerId] != nil && tr.fe[followerId][followeeId] != nil {
		return
	}
	if tr.fl[followerId] == nil {
		tr.fl[followerId] = &list.List{}
	}
	tr.fl[followerId].PushBack(followeeId)
	if tr.fe[followerId] == nil {
		tr.fe[followerId] = make(map[int]*list.Element)
	}
	tr.fe[followerId][followeeId] = tr.fl[followerId].Back()
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (tr *Twitter) Unfollow(followerId int, followeeId int) {
	if tr.fe[followerId] == nil || tr.fe[followerId][followeeId] == nil {
		return
	}
	e := tr.fe[followerId][followeeId]
	tr.fe[followerId][followeeId] = nil
	tr.fl[followerId].Remove(e)
}

// 这里是为了实现合并
type intHp []*node

func (ih *intHp) Len() int {
	return len(*ih)
}

func (ih *intHp) Less(i, j int) bool {
	return (*ih)[i].weight < (*ih)[j].weight
}

func (ih *intHp) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

func (ih intHp) Top() *node {
	if ih.Len() == 0 {
		return nil
	}
	return ih[0]
}

func (ih *intHp) Pop() interface{} {
	if ih.Len() == 0 {
		return -1
	}
	bc := (*ih)[ih.Len()-1]
	*ih = (*ih)[:ih.Len()-1]
	return bc
}

func (ih *intHp) Push(bc interface{}) {
	*ih = append(*ih, bc.(*node))
}
