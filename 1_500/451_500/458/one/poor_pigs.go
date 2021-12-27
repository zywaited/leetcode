package one

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// minutesToTest时间能验证的bucket数
	base := minutesToTest/minutesToDie + 1
	ans := 0
	ab := 1
	for ab < buckets {
		// 可能会有超过int限制
		ab *= base
		ans++
	}
	return ans
}
