package one

func Merge(intervals [][]int) [][]int {
	il := len(intervals)
	if il < 2 {
		return intervals
	}

	type node struct {
		rn   []int
		next *node
	}

	var (
		r        [][]int // 返回结果
		interval []int   // 当前遍历区间
		cn       *node   // 当前节点
		mg       bool    // 是否已经更新过
		change   bool    // 是否区间发生改变
		prev     *node   // 前一个节点
		up       *node   // 当前需要更新区间的节点
	)
	reduce := &node{rn: intervals[0]}
	for i := 1; i < il; i++ {
		cn = reduce
		up = reduce
		prev = nil
		mg = false
		change = false
		interval = intervals[i]
		for {
			// 是否需要合并
			if interval[0] <= up.rn[1] && interval[1] >= up.rn[0] {
				if interval[0] < up.rn[0] {
					up.rn[0] = interval[0]
					change = true
				}
				if interval[1] > up.rn[1] {
					up.rn[1] = interval[1]
					change = true
				}
				if mg && prev != nil {
					// 需要删除
					prev.next = cn.next
				} else {
					prev = cn
					mg = true
				}

				if !change || cn.next == nil {
					break
				}
				cn = cn.next
				interval = cn.rn
				continue
			}
			// 无需再处理
			if mg {
				break
			}

			// 小
			if interval[1] < cn.rn[0] {
				if prev != nil {
					prev.next = &node{rn: intervals[i], next: cn}
				} else {
					reduce = &node{rn: intervals[i], next: cn}
				}
				break
			}
			// 尾
			if cn.next == nil {
				cn.next = &node{rn: intervals[i]}
				break
			}
			prev = cn
			cn = cn.next
			up = cn
		}
	}
	for reduce != nil {
		r = append(r, reduce.rn)
		reduce = reduce.next
	}
	return r
}
