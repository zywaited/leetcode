package one

// sumA-x+y = sumB+x-y
// sumA-sumB = 2(x-y)
// x = (sumA-sumB)/2+y
func FairCandySwap(A []int, B []int) []int {
	sumA := 0
	am := make(map[int]bool)
	for _, num := range A {
		sumA += num
		am[num] = true
	}
	sumB := 0
	for _, num := range B {
		sumB += num
	}
	diff := (sumA - sumB) >> 1
	for _, num := range B {
		if am[diff+num] {
			return []int{diff + num, num}
		}
	}
	return nil
}
