package one

func SortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// 先对组间排序
	// 再对组内排序

	tp := func(items, nums []int, dependOn [][]int) []int {
		queue := make([]int, 0, len(items))
		data := make([]int, 0, len(items))
		for _, item := range items {
			if nums[item] == 0 {
				queue = append(queue, item)
			}
		}
		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]
			data = append(data, item)
			for _, nextItem := range dependOn[item] {
				nums[nextItem]--
				if nums[nextItem] == 0 {
					queue = append(queue, nextItem)
				}
			}
		}
		return data
	}

	// 组被依赖情况
	groupDependOn := make([][]int, m+n)
	// 组依赖数量
	groupDependOnNum := make([]int, m+n)
	// 项目被依赖情况
	memberDependOn := make([][]int, n)
	// 项目依赖数量
	memberDependOnNum := make([]int, n)
	// 每个组的项目情况
	groupItems := make([][]int, m+n)
	for itemIndex := range group {
		if group[itemIndex] == -1 {
			group[itemIndex] = m + itemIndex
		}
		groupItems[group[itemIndex]] = append(groupItems[group[itemIndex]], itemIndex)
	}
	for itemIndex, beforeItem := range beforeItems {
		for _, item := range beforeItem {
			if group[itemIndex] != group[item] {
				// 不同组
				groupDependOn[group[item]] = append(groupDependOn[group[item]], group[itemIndex])
				groupDependOnNum[group[itemIndex]]++
			} else {
				memberDependOn[item] = append(memberDependOn[item], itemIndex)
				memberDependOnNum[itemIndex]++
			}
		}
	}

	groups := make([]int, m+n)
	for groupIndex := range groups {
		groups[groupIndex] = groupIndex
	}
	groups = tp(groups, groupDependOnNum, groupDependOn)
	if len(groups) < m+n {
		return nil
	}

	ans := make([]int, 0, n)
	for _, groupIndex := range groups {
		items := tp(groupItems[groupIndex], memberDependOnNum, memberDependOn)
		if len(items) < len(groupItems[groupIndex]) {
			return nil
		}
		ans = append(ans, items...)
	}
	return ans
}
