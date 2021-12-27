package one

/**
 * one 代表出现一次的
 * two 代表出现两次
 * three 代表出现三次
 */
func SingleNumber(nums []int) int {
	one, two, three := 0, 0, 0
	for _, num := range nums {
		// 是否出现两次
		// one永远为当前出现一次的数据，所以two需要或运算
		two |= one & num
		// 因为出现一次则one对应位为1，两次则two对应位为1， 出现三次则one和two对应位都为1
		one ^= num
		// 是否出现三次
		// 每次满三次的当前数据被去掉，所以用等号
		three = one & two
		// 去除出现三次的数据
		one &= ^three
		two &= ^three
	}
	return one
}
