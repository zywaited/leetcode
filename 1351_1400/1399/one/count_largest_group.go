package one

func CountLargestGroup(n int) int {
	// 已经计算过的数位和
	sm := make(map[int]int, n)
	// 数位和的数量
	// 最大和36[9999], 所以长度就只有37, 这里也可以用数组
	nm := make(map[int]int, 37)
	max := 0
	for num := 1; num <= n; num++ {
		sm[num] = sm[num/10] + num%10
		nm[sm[num]]++
		if nm[sm[num]] > max {
			max = nm[sm[num]]
		}
	}
	ans := 0
	for _, num := range nm {
		if num == max {
			ans++
		}
	}
	return ans
}
