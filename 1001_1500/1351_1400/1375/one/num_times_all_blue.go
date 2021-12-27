package one

func NumTimesAllBlue(light []int) int {
	ans := 0
	// 关闭的灯
	cl := make([]int, len(light))
	// 亮灯最大值
	li := len(light)
	// 灭灯的最小值
	ci := len(light) + 1
	for k := len(light) - 1; k >= 0; k-- {
		// k在ci左侧并且ci右侧的灯都关闭
		if light[k] < ci && li < ci {
			ans++
		}
		// 关闭第k个灯
		if light[k] < ci {
			ci = light[k]
		}
		cl[k] = 1
		// 找到最大亮灯值
		if light[k] < li || k == 0 {
			continue
		}
		// 这里整体也就循环了N次(所以整个过程最坏时间复杂度为O(2N))
		for ; li > 0; li-- {
			if cl[li-1] == 0 {
				break
			}
		}
	}
	return ans
}
