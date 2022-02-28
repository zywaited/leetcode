package one

func MaximumRequests(n int, requests [][]int) int {
	ans := 0
	nums := make([]int, n)
	for i := 1; i < 1<<uint(len(requests)); i++ {
		num := 0
		for j := 0; j < len(requests); j++ {
			if (i>>j)&1 == 0 {
				continue
			}
			num++
			nums[requests[j][0]]--
			nums[requests[j][1]]++
		}
		for j := 0; j < n; j++ {
			if nums[j] != 0 {
				num = 0
				nums[j] = 0
			}
		}
		if ans < num {
			ans = num
		}
	}
	return ans
}
