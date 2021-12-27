package one

func FindIntegers(n int) int {
	length := 1
	for (1 << uint(length)) <= n {
		length++
	}
	dp := make([][]int, length+1)
	dp[0] = []int{1, 0}
	for i := 1; i <= length; i++ {
		dp[i] = []int{dp[i-1][0] + dp[i-1][1], dp[i-1][0]}
	}
	num := 0
	prev := false
	for i := 1; i <= length; i++ {
		if (1<<uint(length-i))&n == 0 {
			prev = false
			continue
		}
		num += dp[length-i][0] + dp[length-i][1]
		if prev {
			return num
		}
		prev = true
	}
	return num + 1
}
