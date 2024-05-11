package one

func garbageCollection(garbage []string, travel []int) int {
	cs := [3]int{}
	ms := 0
	ts := 0
	for i := range garbage {
		ms += len(garbage[i])
		if i > 0 {
			ts += travel[i-1]
		}
		for j := range garbage[i] {
			switch garbage[i][j] {
			case 'M':
				cs[0] = ts
			case 'P':
				cs[1] = ts
			case 'G':
				cs[2] = ts
			}
		}
	}
	for _, m := range cs {
		ms += m
	}
	return ms
}
