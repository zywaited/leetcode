package two

func NumSubarraysWithSum(nums []int, goal int) int {
	sums := map[int]int{0: 1}
	sum := 0
	ans := 0
	for _, num := range nums {
		sum += num
		ans += sums[sum-goal]
		sums[sum]++
	}
	return ans
}
