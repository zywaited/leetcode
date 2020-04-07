package two

func NumTimesAllBlue(light []int) int {
	ans := 0
	// 亮灯最大位置
	li := 0
	for k := range light {
		// 获取最大位置
		if light[k] > li {
			li = light[k]
		}
		// 最大位置就是当前，那么刚好满足条件
		if li == k+1 {
			ans++
		}
	}
	return ans
}
