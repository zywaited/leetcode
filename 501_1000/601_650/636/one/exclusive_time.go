package one

import (
	"strconv"
	"strings"
)

func ExclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	fns := make([]int, 0, n)
	curr := 0
	for _, log := range logs {
		info := strings.Split(log, ":")
		id, _ := strconv.Atoi(info[0])
		ct, _ := strconv.Atoi(info[2])
		switch info[1] {
		case "start":
			if len(fns) > 0 {
				ans[fns[len(fns)-1]] += ct - curr
			}
			fns = append(fns, id)
			curr = ct
		case "end":
			ans[id] += ct - curr + 1
			curr = ct + 1
			fns = fns[:len(fns)-1]
		}
	}
	return ans
}
