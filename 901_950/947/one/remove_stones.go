package one

func RemoveStones(stones [][]int) int {
	group := make([][]int, len(stones))
	for i, s := range stones {
		for j, ns := range stones {
			if s[0] == ns[0] || s[1] == ns[1] {
				group[i] = append(group[i], j)
			}
		}
	}
	// 相连的保留一个即可
	used := make([]bool, len(stones))
	var dfs func(int)
	dfs = func(i int) {
		used[i] = true
		for _, j := range group[i] {
			if !used[j] {
				dfs(j)
			}
		}
	}
	cnt := 0
	for i := range used {
		if !used[i] {
			cnt++
			dfs(i)
		}
	}
	return len(stones) - cnt
}
