package one

func IsMagic(target []int) bool {
	m := make([]int, 0, len(target))
	// 第一次处理
	for n := 2; n <= len(target); n += 2 {
		m = append(m, n)
	}
	for n := 1; n <= len(target); n += 2 {
		m = append(m, n)
	}
	k := 0
	for k < len(target) && target[k] == m[k] {
		k++
	}
	if k == 0 {
		return false
	}
	m = m[k:]
	target = target[k:]
	tm := make([]int, 0, len(target))
	index := 0
	fetchK := 0
	// 模拟过程
	for len(m) > 0 {
		tm = tm[:0]
		for index = 1; index < len(m); index += 2 {
			tm = append(tm, m[index])
		}
		for index = 0; index < len(m); index += 2 {
			tm = append(tm, m[index])
		}
		for fetchK = 0; fetchK < k && fetchK < len(tm); fetchK++ {
			if tm[fetchK] != target[fetchK] {
				return false
			}
		}
		if len(tm) < k {
			break
		}
		m, tm = tm, m
		m = m[k:]
		target = target[k:]
	}
	return true
}
