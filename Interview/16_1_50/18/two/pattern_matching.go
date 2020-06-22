package two

func PatternMatching(pattern string, value string) bool {
	// 把方法一改为循环
	ac := 0
	bc := 0
	for i := range pattern {
		if pattern[i] == 'a' {
			ac++
			continue
		}
		bc++
	}
	// 始终以数量多的字符为中心计算长度
	if ac < bc {
		ac, bc = bc, ac
		tmp := make([]byte, 0, len(pattern))
		for i := range pattern {
			if pattern[i] == 'a' {
				tmp = append(tmp, 'b')
				continue
			}
			tmp = append(tmp, 'a')
		}
		pattern = string(tmp)
	}
	if len(value) == 0 {
		return bc == 0 // 只能保留一个符号
	}
	if len(pattern) == 0 {
		return len(value) == 0
	}
	// a的长度
	al := 0
	for ; al*ac <= len(value); al++ {
		tbl := len(value) - al*ac
		if (bc == 0 && tbl == 0) || (bc > 0 && tbl%bc == 0) {
			// 剩余字符能够完整拼接b数据
			bl := 0
			if bc > 0 {
				bl = tbl / bc
			}
			as, bs := "", ""
			pos := 0     // value index
			flag := true // 是否匹配完成
			for i := range pattern {
				if pattern[i] == 'a' {
					sub := value[pos : pos+al]
					if len(as) == 0 {
						as = sub
					} else if as != sub {
						flag = false
						break
					}
					pos += al
					continue
				}
				sub := value[pos : pos+bl]
				if len(bs) == 0 {
					bs = sub
				} else if bs != sub {
					flag = false
					break
				}
				pos += bl
			}
			if flag && as != bs {
				return true
			}
		}
	}
	return false
}
