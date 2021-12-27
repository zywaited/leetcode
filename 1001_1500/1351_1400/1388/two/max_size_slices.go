package two

import "container/list"

// 贪心(反悔)+双向列表+递减链表(类似优先队列))
func MaxSizeSlices(slices []int) int {
	// 数组转换成链表方便删除
	ll := make([]int, len(slices))
	lr := make([]int, len(slices))
	// 索引是否被删除
	rm := make(map[int]byte, len(slices))
	// 本来是使用优先队列(二叉树),没有现成的包,简单处理下(量不大)
	ls := &list.List{}
	add := func(index int) {
		if ls.Len() == 0 {
			ls.PushFront(index)
			return
		}
		cl := ls.Front()
		for ; cl != nil && slices[cl.Value.(int)] > slices[index]; cl = cl.Next() {
		}
		if cl == nil {
			ls.PushBack(index)
			return
		}
		ls.InsertBefore(index, cl)
	}
	for index := range slices {
		ll[index] = index - 1
		if index == 0 {
			ll[index] = len(slices) - 1
		}
		lr[index] = index + 1
		if index == len(slices)-1 {
			lr[index] = 0
		}
		add(index)
	}
	ans := 0
	mj := len(slices) / 3 // 最大选择次数
	for ; mj > 0; mj-- {
		cl := ls.Front()
		for rm[cl.Value.(int)] == 1 {
			cl = cl.Next()
			// 删除
			ls.Remove(cl.Prev())
		}
		index := cl.Value.(int)
		ans += slices[index]
		// 当前值重置为差值
		ls.Remove(cl)
		slices[index] = slices[ll[index]] + slices[lr[index]] - slices[index]
		// 左右变为删除状态
		rm[ll[index]] = 1
		rm[lr[index]] = 1
		// 重新放入数据
		add(index)
		// 更新当前的左右节点
		ll[index] = ll[ll[index]]
		lr[index] = lr[lr[index]]
		// 更新左右节点
		lr[ll[ll[index]]] = index
		ll[lr[lr[index]]] = index
	}
	return ans
}
