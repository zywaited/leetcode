package one

func LongestConsecutive(nums []int) int {
	mp := make(map[int]byte)
	// 先把数字都放入map中
	for _, num := range nums {
		mp[num] = 0
	}

	max := 0
	for _, num := range nums {
		// 为了避免重复计算
		if _, ok := mp[num-1]; ok {
			continue
		}
		cn := num
		n := 1
		for {
			// 循环找连续数字
			if _, ok := mp[cn+1]; !ok {
				break
			}
			cn++
			n++
		}
		if n > max {
			max = n
		}
	}
	return max
}
