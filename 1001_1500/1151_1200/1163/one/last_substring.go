package one

// 暴力
func LastSubstring(s string) string {
	indexes := make([]int, 0)
	max := byte(0)
	// 先找出最da的字符
	for si := range s {
		if s[si] == max {
			indexes = append(indexes, si)
			continue
		}
		if s[si] > max {
			max = s[si]
			indexes = indexes[:0]
			indexes = append(indexes, si)
		}
	}
	// 如果只有一个(取巧的做法)
	if len(indexes) == len(s) {
		return s
	}
	// 两两比较
	i := compare(s, indexes)
	if i == -1 {
		return "" // 理论上不会走到这里
	}
	return s[i:]
}

func compare(s string, index []int) int {
	num := len(index)
	if num == 0 {
		return -1
	}
	if num == 1 {
		return index[0]
	}
	fi, si := -1, -1
	if num > 2 {
		middle := num >> 1
		fi, si = compare(s, index[:middle]), compare(s, index[middle:])
	} else {
		fi, si = index[0], index[1]
	}
	if fi == -1 {
		return si
	}
	if si == -1 {
		return fi
	}
	i, j := fi, si
	for {
		if i < len(s) && j < len(s) {
			if s[i] > s[j] {
				return fi
			}
			if s[i] < s[j] {
				return si
			}
			i++
			j++
			continue
		}
		if i == len(s) {
			return si
		}
		return fi
	}
}
