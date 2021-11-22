package two

func NumFriendRequests(ages []int) int {
	// 每个年龄的人数
	an := [121]int{}
	for _, age := range ages {
		an[age]++
	}
	// 前缀和
	ps := [121]int{}
	for i := 1; i < len(an); i++ {
		ps[i] = ps[i-1] + an[i]
	}
	ans := 0
	for age := 1; age < len(an); age++ {
		if an[age] > 0 && (age>>1)+7 < age {
			ans += (ps[age] - ps[(age>>1)+7] - 1) * an[age]
		}
	}
	return ans
}
