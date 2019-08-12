package one

func RemoveElement(nums []int, val int) int {
	rs := 0
	for _, num := range nums {
		if num^val != 0 {
			nums[rs] = num
			rs++
		}
	}
	return rs
}
