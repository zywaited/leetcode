package one

func PredictPartyVictory(senate string) string {
	st := []byte(senate)
	r, d := 0, 0
	for {
		i := 0
		for _, s := range st {
			if s == 'R' {
				if r > 0 {
					r--
					continue
				}
				d++
			} else {
				if d > 0 {
					d--
					continue
				}
				r++
			}
			st[i] = s
			i++
		}
		if i == len(st) {
			// 没变化
			break
		}
		st = st[:i]
	}
	if st[0] == 'R' {
		return "Radiant"
	}
	return "Dire"
}
