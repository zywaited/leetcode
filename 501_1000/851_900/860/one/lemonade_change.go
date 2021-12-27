package one

func LemonadeChange(bills []int) bool {
	n5, n10 := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			n5++
		case 10:
			n10++
			n5--
		case 20:
			if n10 == 0 {
				n5 -= 3
			} else {
				n10--
				n5--
			}
		}
		if n5 < 0 || n10 < 0 {
			return false
		}
	}
	return true
}
