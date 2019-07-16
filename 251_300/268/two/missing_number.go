package two

func MissingNumber(nums []int) int {
	// 其实就是找只出现一次的数字
	// 自身异或为0
	r := len(nums)
	for i, num := range nums {
		r ^= i
		r ^= num
	}
	return r
}
