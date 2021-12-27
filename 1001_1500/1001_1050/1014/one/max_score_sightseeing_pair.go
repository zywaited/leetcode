package one

func MaxScoreSightseeingPair(A []int) int {
	ans := 0
	pi := 0
	for i := 1; i < len(A); i++ {
		ans = max(ans, A[pi]+A[i]+pi-i)
		if A[pi]-A[i] <= i-pi { // 这种情况距离拉开了，实际分值变小，所以变更点
			pi = i
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
