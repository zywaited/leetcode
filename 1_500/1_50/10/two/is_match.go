package two

func IsMatch(s string, p string) bool {
	// 循环版
	sl := len(s)
	pl := len(p)
	var (
		asIndex [][2]int
		asl     int
		skipAs  int
	)
	i, j := 0, 0
	for i < sl {
		// 正常匹配或者是通配符匹配
		if j < pl && (s[i] == p[j] || p[j] == '.') {
			// *匹配
			if j < pl-1 && p[j+1] == '*' {
				// 保留当前匹配位置
				asIndex = append(asIndex, [2]int{j, i - 1})
				// 跳过*，下一位
				j += 2
				continue
			}
			i++
			j++
			continue
		}

		// *匹配
		if j < pl-1 && p[j+1] == '*' {
			// 跳过*，下一位
			skipAs++
			j += 2
			continue
		}

		// 匹配失败查看是否可以回到*保留的位置
		for {
			if len(asIndex) == 0 {
				return false
			}
			asl = len(asIndex) - 1
			// *号位置继续匹配
			j = asIndex[asl][0] + 2
			// *匹配所有，所以此处加1继续匹配下一个
			asIndex[asl][1]++
			// 判断是否与*前第一个字符匹配
			if s[asIndex[asl][1]] != p[asIndex[asl][0]] && p[asIndex[asl][0]] != '.' {
				// 1 为什么这里要小于0, 因为即便匹配不成功，但是*可以匹配空，去除前一个字符
				// 2 为什么要这个判断，其实没有这个判断程序也是正常的
				// 不过既然都是正常匹配，前一个回溯失败，回溯多个也会失败
				// 因为这里skip减1如果未走到skip加1逻辑，代表两个*中间一定有字符无法被匹配
				if skipAs < 0 {
					return false
				}
				skipAs--
				asIndex = asIndex[:asl]
				continue
			}
			i = asIndex[asl][1] + 1
			break
		}
	}

	if j >= pl {
		return true
	}

	// s匹配完成后，如果p还有字符就只能是*或者字符加*组合
	p = p[j:]
	pl = len(p)
	j = 0
	for j < pl {
		if p[j] == '*' {
			j++
			continue
		}
		if j < pl-1 && p[j+1] == '*' {
			j += 2
			continue
		}
		return false
	}
	return true
}
