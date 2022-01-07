package one

func MaxVacationDays(flights [][]int, days [][]int) int {
	// Write your code here
	n := len(flights)
	k := len(days[0])
	curr := make([]int, n)
	for i := range curr {
		curr[i] = -1
	}
	curr[0] = 0
	next := make([]int, n)
	ans := 0
	for tk := 1; tk <= k; tk++ {
		for i := 0; i < n; i++ {
			next[i] = -1
			for j := 0; j < n; j++ {
				if (i != j && flights[j][i] == 0) || curr[j] < 0 {
					continue
				}
				if next[i] < curr[j]+days[i][tk-1] {
					next[i] = curr[j] + days[i][tk-1]
				}
				if tk == k && ans < next[i] {
					ans = next[i]
				}
			}
		}
		curr, next = next, curr
	}
	return ans
}
