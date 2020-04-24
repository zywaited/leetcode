package one

// 归并排序
func ReversePairs(nums []int) int {
	return ms(nums, 0, len(nums)-1)
}

func ms(nums []int, l, r int) int {
	// 边界值
	if l >= r {
		return 0
	}
	// 与快排不同的是以中心为基础拆分
	// 这样刚好就可以计算左右区间的逆序对，然后再合并计算
	m := (l + r) >> 1
	num := ms(nums, l, m) + ms(nums, m+1, r)
	tmp := make([]int, r-l+1) // 临时存储排序数据
	s := l
	e := m + 1
	i := 0
	for ; s <= m && e <= r; i++ {
		if nums[s] <= nums[e] {
			// 左边比右边小
			// 这时看右边偏移量就是当前s的逆序对
			tmp[i] = nums[s]
			s++
			num += e - m - 1
			continue
		}
		// 右边不会有逆序对(排好序了)
		tmp[i] = nums[e]
		e++
	}
	// 对最后剩余的数据做处理
	for ; s <= m; s++ {
		// 左边区间还有
		tmp[i] = nums[s]
		i++
		num += e - m - 1
	}
	for ; e <= r; e++ {
		tmp[i] = nums[e]
		i++
	}
	// 处理原数据
	for i = range tmp {
		nums[l+i] = tmp[i]
	}
	return num
}
