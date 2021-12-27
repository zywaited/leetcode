package one

import "sort"

func AccountsMerge(accounts [][]string) [][]string {
	em := make(map[string]int, len(accounts))
	indexes := make([]int, len(accounts)+1)
	for index := range indexes {
		indexes[index] = index
	}
	find := func(index int) int {
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return index
	}
	union := func(f, s int) {
		indexes[find(f)] = indexes[find(s)]
	}
	for index, account := range accounts {
		for emailIndex := 1; emailIndex < len(account); emailIndex++ {
			if em[account[emailIndex]] > 0 {
				union(index+1, em[account[emailIndex]])
			}
			em[account[emailIndex]] = find(index + 1)
		}
	}
	ans := make([][]string, 0, len(accounts))
	am := make([]int, len(accounts)+1)
	for ansIndex := range am {
		am[ansIndex] = -1
	}
	for index, account := range accounts {
		accountIndex := find(index + 1)
		if am[accountIndex] == -1 {
			ans = append(ans, []string{account[0]})
			am[accountIndex] = len(ans) - 1

		}
		for emailIndex := 1; emailIndex < len(account); emailIndex++ {
			if em[account[emailIndex]] != -1 {
				ans[am[accountIndex]] = append(ans[am[accountIndex]], account[emailIndex])
				em[account[emailIndex]] = -1
			}
		}
	}
	for index := range ans {
		sort.Strings(ans[index][1:])
	}
	return ans
}
