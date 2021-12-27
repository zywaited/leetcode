package one

func SingleNumber(nums []int) int {
	r := 0
	for _, num := range nums {
		r ^= num
	}
	return r
}
