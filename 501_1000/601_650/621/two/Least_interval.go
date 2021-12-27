package two

func LeastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for _, t := range tasks {
		cnt[t-'A']++
	}
	mt, mn := 0, 0
	// 以数量最多的为主
	for _, t := range cnt {
		if t > mt {
			mt = t
			mn = 1
			continue
		}
		if t == mt {
			mn++
		}
	}
	ct := (mt-1)*(n+1) + mn
	if ct < len(tasks) {
		ct = len(tasks)
	}
	return ct
}
