package one

func FourSumCount(A []int, B []int, C []int, D []int) int {
	ab := make(map[int]int)
	for _, an := range A {
		for _, bn := range B {
			ab[an+bn]++
		}
	}
	n := 0
	for _, cn := range C {
		for _, dn := range D {
			n += ab[0-(cn+dn)]
		}
	}
	return n
}
