package one

const mod = 1e9 + 7

func NumPrimeArrangements(n int) int {
	// 由于题目已经给了n的范围，偷个懒就不用使用程序查找对应的质数个数
	// 如果需要程序计算，请参考204题: https://leetcode-cn.com/problems/count-primes/
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	// 质数个数
	nums := 0
	for _, num := range primes {
		if num <= n {
			nums++
			continue
		}
		break
	}
	return int((count(nums) * count(n-nums)) % mod)
}

// 获取排列个数A(N,N)
func count(n int) int64 {
	// 可能会超32位
	count := int64(1)
	for ; n > 0; n-- {
		count = (count * int64(n)) % mod
	}
	return count
}
