package one

func MaxOperations(nums []int, k int) int {
	nm := make(map[int]int, len(nums))
	t := 0
	for _, n := range nums {
		if nm[k-n] > 0 {
			t++
			nm[k-n]--
			continue
		}
		nm[n]++
	}
	return t
}
