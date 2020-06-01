package one

func KidsWithCandies(candies []int, extraCandies int) []bool {
	// 先找到最大值
	max := 0
	for _, num := range candies {
		if num > max {
			max = num
		}
	}
	ans := make([]bool, len(candies))
	for index, num := range candies {
		if num+extraCandies >= max {
			ans[index] = true
		}
	}
	return ans
}
