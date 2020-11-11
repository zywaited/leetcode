package one

func FindRotateSteps(ring string, key string) int {
	// 求两点最短距离
	indexes := make([][]int, 26)
	for index := range ring {
		indexes[ring[index]-'a'] = append(indexes[ring[index]-'a'], index)
	}

	steps := make([][]int, len(ring))
	var sp func(int, int) int
	sp = func(rIndex int, kIndex int) int {
		if len(steps[rIndex]) == 0 {
			steps[rIndex] = make([]int, len(key))
		}
		if kIndex == len(key) {
			return 0
		}
		if steps[rIndex][kIndex] > 0 {
			return steps[rIndex][kIndex]
		}
		step := len(key) * len(key) // math.MaxInt32
		for _, index := range indexes[key[kIndex]-'a'] {
			cm := 0
			if rIndex < index {
				cm = min(index-rIndex, len(ring)-index+rIndex)
			}
			if rIndex > index {
				cm = min(rIndex-index, len(ring)-rIndex+index)
			}
			step = min(step, sp(index, kIndex+1)+cm)
		}
		step++ // 书写
		steps[rIndex][kIndex] = step
		return step
	}
	return sp(0, 0)
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
