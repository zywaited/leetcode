package one

func MaximalNetworkRank(n int, roads [][]int) int {
	nums := [101]int{}
	connects := make(map[int]struct{}, n<<1)
	for _, road := range roads {
		nums[road[0]]++
		nums[road[1]]++
		connects[road[0]*n+road[1]] = struct{}{}
		connects[road[1]*n+road[0]] = struct{}{}
	}
	ans := 0
	num := 0
	ok := true
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			num = nums[i] + nums[j]
			if _, ok = connects[i*n+j]; ok {
				num--
			}
			ans = max(ans, num)
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
