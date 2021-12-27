package one

// 空间复杂度: O(N)
func ProductExceptSelf(nums []int) []int {
	nl := len(nums)
	prefix := make([]int, nl) // 前乘积
	prefix[0] = 1
	i := 0
	for i = 1; i < nl; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	suffix := make([]int, nl) // 后乘积
	suffix[nl-1] = 1
	for i = nl - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	rs := make([]int, nl)
	for i = range rs {
		rs[i] = prefix[i] * suffix[i]
	}
	return rs
}
