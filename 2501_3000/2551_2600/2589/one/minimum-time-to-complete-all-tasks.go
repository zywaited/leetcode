package one

import "sort"

func findMinimumTime(tasks [][]int) int {
	used := [2001]bool{}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	costTime := 0
	for _, task := range tasks {
		duration := task[2]
		for i := task[0]; i <= task[1] && duration > 0; i++ {
			if used[i] {
				duration--
			}
		}
		if duration <= 0 {
			continue
		}
		for i := task[1]; duration > 0; i-- {
			if !used[i] {
				duration--
				costTime++
				used[i] = true
			}
		}
	}
	return costTime
}
