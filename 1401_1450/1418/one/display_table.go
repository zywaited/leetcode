package one

import (
	"sort"
	"strconv"
)

func DisplayTable(orders [][]string) [][]string {
	m := make(map[string]bool, len(orders)) // 名称没有数字，所以直接公用
	ts := make([]string, 1, len(orders))    // 表头，预留一个给 Table
	ts[0] = "Table"
	os := make([]int, 0, len(orders))                  // 桌号
	nm := make(map[string]map[string]int, len(orders)) // 餐名对应桌号的数量
	for _, order := range orders {
		if !m[order[1]] {
			m[order[1]] = true
			n, _ := strconv.Atoi(order[1])
			os = append(os, n)
		}
		if !m[order[2]] {
			m[order[2]] = true
			ts = append(ts, order[2])
		}
		if nm[order[2]] == nil {
			nm[order[2]] = make(map[string]int)
		}
		nm[order[2]][order[1]]++
	}
	sort.Strings(ts[1:])
	sort.Ints(os)
	ans := make([][]string, len(os)+1)
	ans[0] = ts
	for i := 1; i < len(ans); i++ {
		ans[i] = make([]string, len(ts))
		ans[i][0] = strconv.Itoa(os[i-1])
		for j := 1; j < len(ts); j++ {
			if nm[ts[j]] == nil {
				ans[i][j] = "0"
				continue
			}
			ans[i][j] = strconv.Itoa(nm[ts[j]][ans[i][0]])
		}
	}
	return ans
}
