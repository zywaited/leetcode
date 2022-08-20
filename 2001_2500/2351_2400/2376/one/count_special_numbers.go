package one

func CountSpecialNumbers(n int) int {
	ns := make([]int, 0)
	for cn := n; cn > 0; cn /= 10 {
		ns = append(ns, cn%10)
	}
	ans := 0
	// less than n
	for cn := 1; cn < len(ns); cn++ {
		tn := 9
		for ci := 1; ci < cn; ci++ {
			tn *= 10 - ci
		}
		ans += tn
	}
	i := len(ns) - 1
	exist := [10]bool{}
	for ; i >= 0; i-- {
		// every num reduce
		if i < len(ns)-1 || ns[i] > 1 {
			tn := 0
			if i == len(ns)-1 {
				tn = ns[i] - 1
			} else {
				for cn := ns[i] - 1; cn >= 0; cn-- {
					if !exist[cn] {
						tn++
					}
				}
			}
			for ci := i; tn > 0 && ci > 0; ci-- {
				tn *= 10 - (len(ns) - ci)
			}
			ans += tn
		}
		if exist[ns[i]] {
			break
		}
		exist[ns[i]] = true
		if i == 0 {
			// self
			ans++
		}
	}
	return ans
}
