package one

import "math"

// 只会暴力枚举
func SumFourDivisors(nums []int) int {
	ans := 0
	add := func(num int) {
		sum := 0
		times := 0
		for i := int(math.Sqrt(float64(num))); i > 0; i-- {
			if num%i > 0 {
				continue
			}
			times++
			if times > 2 || num/i == i {
				// 完全平方数
				return
			}
			sum += i
			sum += num / i
		}
		if times == 2 {
			ans += sum
		}
	}
	for _, num := range nums {
		add(num)
	}
	return ans
}
