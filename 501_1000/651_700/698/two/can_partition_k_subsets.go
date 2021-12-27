package two

func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 每个子集的总和
	tr, ex := sum/k, sum%k
	if ex != 0 {
		return false
	}
	// 所有的集合
	dp := make(map[int]bool, 1<<uint(len(nums)))
	var sr func(int, int) bool
	sr = func(now int, sum int) bool {
		if _, ok := dp[now]; !ok {
			// 所有数字已使用
			if sum == 0 {
				dp[now] = true
				return true
			}
			dp[now] = false // 默认都是无效
			// 当前能够取的最大值(保证分组)
			max := (sum-1)%tr + 1
			for i, num := range nums {
				if (1<<uint(i)&now) != 0 || num > max || !sr(now|(1<<uint(i)), sum-num) {
					continue
				}
				dp[now] = true
				break
			}
		}
		return dp[now]
	}
	return sr(0, sum)
}
