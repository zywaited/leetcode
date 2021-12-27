package one

func HammingWeight(num uint32) int {
	n := 0
	for num != 0 {
		n++
		num &= num - 1
	}
	return n
}
