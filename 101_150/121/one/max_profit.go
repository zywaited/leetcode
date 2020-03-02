package one

func MaxProfit(prices []int) int {
	pn := -1
	ans := 0
	for _, num := range prices {
		if pn == -1 || num <= pn {
			pn = num
			continue
		}
		if num-pn > ans {
			ans = num - pn
		}
	}
	return ans
}
