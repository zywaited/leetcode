package one

func DistributeCandies(candies []int) int {
	// 先算出种类
	// 则答案就是MIN(种类数量, 糖果数量的二分之一)
	nums := make(map[int]byte)
	for _, candy := range candies {
		nums[candy] = 1
	}
	ans := len(candies) >> 1
	if len(nums) < ans {
		ans = len(nums)
	}
	return ans
}
