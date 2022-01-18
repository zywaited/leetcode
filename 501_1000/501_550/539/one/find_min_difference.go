package one

import "sort"

func FindMinDifference(timePoints []string) int {
	points := make([]int, 0, len(timePoints))
	for _, timePoint := range timePoints {
		points = append(points, (int(timePoint[0]-'0')*10+int(timePoint[1]-'0'))*60+int(timePoint[3]-'0')*10+int(timePoint[4]-'0'))
	}
	sort.Ints(points)
	ans := 24*60 - points[len(points)-1] + points[0]
	for i := 1; i < len(points); i++ {
		ans = min(ans, points[i]-points[i-1])
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
