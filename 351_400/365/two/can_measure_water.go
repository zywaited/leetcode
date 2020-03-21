package one

func CanMeasureWater(x int, y int, z int) bool {
	// 不能满足
	if x+y < z {
		return false
	}
	// 去除算公约数的问题
	if x == 0 {
		return x+y == z
	}
	// 灌水一定会得到x,y的最大公约数升水，就判断z是否能够整除就行
	return z%gcd(x, y) == 0
}

func gcd(x, y int) int {
	for y > 0 {
		x, y = y, x%y
	}
	return x
}
