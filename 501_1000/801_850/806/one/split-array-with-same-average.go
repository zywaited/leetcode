package one

func splitArraySameAverage(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	n := len(nums)
	l := n / 2
	isPossible := false
	for i := 1; i <= l; i++ {
		if sum*i%n == 0 {
			isPossible = true
			break
		}
	}
	if !isPossible {
		return false
	}
	dp := make([]map[int]bool, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := min(i, l-1); j >= 0; j-- {
			dl := j + 1
			for prevNum := range dp[j] {
				target := prevNum + num
				dp[dl][target] = true
				if sum*dl == target*n {
					return true
				}
			}
		}
	}
	return false
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
