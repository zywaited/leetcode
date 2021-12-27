package one

func ShipWithinDays(weights []int, D int) int {
	r := 0
	l := 0
	for _, weight := range weights {
		r += weight
		if weight > l {
			l = weight
		}
	}
	for l < r {
		m := l + (r-l)>>1
		if checkShipWithinDays(weights, m, D) {
			r = m
			continue
		}
		l = m + 1
	}
	return l
}

func checkShipWithinDays(weights []int, bestWeight, D int) bool {
	d := 1
	pw := 0
	for _, weight := range weights {
		if weight > bestWeight {
			return false
		}
		pw += weight
		if pw <= bestWeight {
			continue
		}
		d++
		pw = weight
		if d > D {
			return false
		}
	}
	return true
}
