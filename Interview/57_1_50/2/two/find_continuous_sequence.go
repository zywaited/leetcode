package two

func FindContinuousSequence(target int) [][]int {
	ans := make([][]int, 0)
	sum := 1
	l, r := 1, 1
	for num := 2; num < target; num++ {
		r = num
		sum += num
		for sum > target {
			sum -= l
			l++
		}
		if sum == target {
			ans = append(ans, make([]int, 0, r-l+1))
			for k := l; k <= r; k++ {
				ans[len(ans)-1] = append(ans[len(ans)-1], k)
			}
		}
	}
	return ans
}
