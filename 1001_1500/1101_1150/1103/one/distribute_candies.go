package one

func DistributeCandies(candies int, people int) []int {
	// 分发轮次
	times := 0
	need := ((people + 1) * people) >> 1
	for candies > need {
		times++
		candies -= need
		need += people * people
	}
	ans := make([]int, people)
	base := 0      // 轮发N次的基础数据
	if times > 1 { // 0或1基础数据都是0（N-1）
		base = (((1+times)*times)>>1 - times) * people
	}
	for i := range ans {
		ans[i] = base + times*(i+1)
		if candies == 0 {
			continue
		}
		// 下一次应该处理的数据
		if candies > (people*times + i + 1) {
			ans[i] += people*times + i + 1
			candies -= people*times + i + 1
			continue
		}
		ans[i] += candies
		candies = 0
	}
	return ans
}
