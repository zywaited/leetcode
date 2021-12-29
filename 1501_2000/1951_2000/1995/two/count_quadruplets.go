package two

func CountQuadruplets(nums []int) int {
	ans := 0
	// 两个数相加和的组合数
	sums := make(map[int]int)
	for i := 1; i < len(nums)-1; i++ {
		for j := i + 1; j > 1 && j < len(nums); j++ {
			ans += sums[nums[j]-nums[i]]
		}
		for j := 0; j < i; j++ {
			sums[nums[j]+nums[i]]++
		}
	}
	return ans
}
