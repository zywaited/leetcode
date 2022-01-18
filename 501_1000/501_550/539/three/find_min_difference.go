package three

func FindMinDifference(timePoints []string) int {
	exists := make([]bool, 24*60)
	for _, timePoint := range timePoints {
		curr := (int(timePoint[0]-'0')*10+int(timePoint[1]-'0'))*60 + int(timePoint[3]-'0')*10 + int(timePoint[4]-'0')
		if exists[curr] {
			return 0
		}
		exists[curr] = true
	}
	ans := len(exists)
	minPoint := -1
	maxPoint := 0
	prev := -1
	for point, exist := range exists {
		if !exist {
			continue
		}
		if minPoint == -1 {
			minPoint = point
		}
		maxPoint = point
		if prev >= 0 {
			ans = min(ans, point-prev)
		}
		prev = point
	}
	ans = min(ans, len(exists)-maxPoint+minPoint)
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
