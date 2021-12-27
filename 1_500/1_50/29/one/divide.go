package one

import "math"

func Divide(dividend int, divisor int) int {
	mark := dividend ^ divisor
	// 特殊判断
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == divisor {
		return 1
	}
	// 溢出
	if divisor == math.MinInt32 {
		return 0
	}
	isMax := false
	if dividend < 0 {
		if dividend == math.MinInt32 {
			dividend = math.MaxInt32
			isMax = true
		} else {
			dividend = -dividend
		}
	}
	if divisor < 0 {
		divisor = -divisor
	}
	resp := 0
	i := 1
	// 优化点：第一次减divisor，下一次就可以尝试减divisor>>1，依次类推
	origin := divisor
	for dividend >= divisor {
		dividend -= divisor
		// 这样减少计算复杂度
		divisor <<= 1
		resp += i
		i <<= 1
		// 还原原始数据
		if dividend < divisor && divisor > origin {
			divisor = origin
			i = 1
			continue
		}
	}
	// 超过最大值，需要此处要加1
	if isMax && dividend+1 >= origin {
		resp++
	}
	if mark < 0 {
		resp = 0 - resp
	}
	return resp
}
