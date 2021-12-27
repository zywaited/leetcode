package two

// 空间复杂度: O(1)
// 先从左到后把前乘积计算出来
// 再从右到左算后乘积
func ProductExceptSelf(nums []int) []int {
	nl := len(nums)
	rs := make([]int, nl)
	rs[0] = 1
	i := 1
	for ; i < nl; i++ {
		rs[i] = rs[i-1] * nums[i-1]
	}
	suffix := 1
	for i = nl - 1; i > 0; i-- {
		rs[i] *= suffix
		suffix *= nums[i]
	}
	rs[0] = suffix
	return rs
}
