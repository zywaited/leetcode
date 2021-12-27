package one

import "math"

func BulbSwitch(n int) int {
	// 因数个数是偶数则该位置最后是关闭的，否则就是打开的
	// 由于每次都会有两个因子相乘，那么数量是奇数，则需要其中一对因子相等
	// 那这样就是找小于等于n有多少个平方数
	return int(math.Sqrt(float64(n)))
}
