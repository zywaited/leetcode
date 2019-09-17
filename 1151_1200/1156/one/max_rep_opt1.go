package one

func MaxRepOpt1(text string) int {
	// 用int而不用byte是为了防止text中有某个字符的数量比byte大
	nums := make(map[int]int) // 每个字符对应的所有数量
	var (
		optNums [][2]int // 每个字符一小段的数量
		cb      int      // 当前字符
		cn      int      // 当前字符长度
		tb      int
	)
	for index := range text {
		tb = int(text[index])
		nums[tb]++
		if index == 0 {
			cb = tb
			cn = 1
			continue
		}
		if tb == cb {
			cn++
			continue
		}
		optNums = append(optNums, [2]int{cb, cn})
		cb = tb
		cn = 1
	}
	optNums = append(optNums, [2]int{cb, cn}) // 最后一个
	m := -1
	for i, optNum := range optNums {
		if optNum[1] > m {
			m = optNum[1]
		}
		// 是否可以替换字符
		if optNum[1] >= nums[optNum[0]] {
			continue
		}
		// 一定可以多一个
		if optNum[1]+1 > m {
			m = optNum[1] + 1
		}
		// 是否替换后后面是连续的
		if i+1 >= len(optNums) || optNums[i+1][1] > 1 ||
			i+2 >= len(optNums) || optNums[i+2][0] != optNum[0] {
			continue
		}
		// 是否是从下一段字符中选的
		if optNum[1]+optNums[i+2][1] >= nums[optNum[0]] {
			if optNum[1]+optNums[i+2][1] > m {
				m = optNum[1] + optNums[i+2][1]
			}
			continue
		}
		if optNum[1]+optNums[i+2][1]+1 > m {
			m = optNum[1] + optNums[i+2][1] + 1
		}
	}
	return m
}
