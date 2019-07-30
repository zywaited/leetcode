package one

// 超时
func FCandy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	mp := make(map[int]int)
	for i := 1; i < len(ratings); i++ {
		// 相等不需要处理
		if ratings[i] == ratings[i-1] {
			continue
		}
		// 如果比上一个大
		if ratings[i] > ratings[i-1] {
			mp[i] = mp[i-1] + 1
			continue
		}
		// 比上一个小但是本来数量就少
		if mp[i] < mp[i-1] {
			continue
		}

		mp[i-1] = mp[i] + 1
		for j := i - 1; j > 0; j-- {
			if ratings[j] >= ratings[j-1] {
				break
			}
			// 本身就大
			if mp[j-1] > mp[j] {
				continue
			}
			mp[j-1] = mp[j] + 1
		}
	}
	rs := len(ratings)
	for _, num := range mp {
		rs += num
	}
	return rs
}

// 优化
func Candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	mp := make(map[int]int)
	for i := 1; i < len(ratings); i++ {
		// 如果比上一个大
		if ratings[i] > ratings[i-1] {
			mp[i] = mp[i-1] + 1
		}
		// 小的后面补偿
	}
	rs := len(ratings) + mp[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		// 大
		if ratings[i] > ratings[i+1] {
			// 小
			if mp[i] <= mp[i+1] {
				mp[i] = mp[i+1] + 1
			}
		}
		rs += mp[i]
	}
	return rs
}
