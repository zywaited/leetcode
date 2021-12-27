package one

import "sort"

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	if len(points) == 0 {
		return 0
	}
	num := 1
	max := points[0][1]
	for _, point := range points {
		if point[0] > max {
			max = point[1]
			num++
		}
	}
	return num
}
