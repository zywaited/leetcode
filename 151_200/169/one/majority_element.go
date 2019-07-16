package one

func MajorityElement(nums []int) int {
	count := 0
	n := 0
	for _, num := range nums {
		if count == 0 || n == num {
			count++
			n = num
			continue
		}
		count--
	}
	// 由于题目给的总是存在众数的，所以后续不需要判断是否当前选出的是否符合
	return n
}
