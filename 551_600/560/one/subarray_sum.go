package one

func SubarraySum(nums []int, k int) int {
	ans := 0
	nm := make(map[int]int, len(nums))
	nm[0] = 1
	ps := 0
	for _, num := range nums {
		ps += num
		// 是否已经存在前缀和(因为用的前缀和，所以一定是连续的)
		if nm[ps-k] > 0 {
			ans += nm[ps-k]
		}
		nm[ps]++
	}
	return ans
}
