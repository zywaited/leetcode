package one

func MinCount(coins []int) int {
	ans := 0
	for _, coin := range coins {
		ans += coin >> 1
		ans += coin & 1
	}
	return ans
}
