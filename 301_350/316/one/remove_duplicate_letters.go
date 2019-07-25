package one

func RemoveDuplicateLetters(s string) string {
	// 空间换时间
	type mp struct {
		data  []byte        // 字符列表
		exist map[byte]bool // 字符是否应存在
	}

	// 计数器
	count := make(map[byte]uint)
	var (
		lastIndex = 0
		i         = 0
	)
	for i = 0; i < len(s); i++ {
		count[s[i]]++
	}

	sk := mp{exist: make(map[byte]bool, len(count))}
	for i = 0; i < len(s); i++ {
		// 如果已经存在
		if sk.exist[s[i]] {
			count[s[i]]--
			continue
		}

		// 字典序
		// 与最后一个字符进行比较(前提是这个字符可以被去掉，也就是count不为0)
		// 因为后面添加的字符一定与前面字符比较过了，所以最后一个字符要比前面字符小或者前面是唯一的了
		for {
			lastIndex = len(sk.data) - 1
			if lastIndex >= 0 && s[i] < sk.data[lastIndex] && count[sk.data[lastIndex]] > 0 {
				// 去除最后一个字符，有更小字典序
				sk.exist[sk.data[lastIndex]] = false
				sk.data = sk.data[:lastIndex]
				continue
			}
			// 加入
			sk.data = append(sk.data, s[i])
			sk.exist[s[i]] = true
			count[s[i]]--
			break
		}
	}
	return string(sk.data)
}
