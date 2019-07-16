package one

/**
 * 先把数组所有的数字进行异或处理，因为有两个数不同，所以一定不会为0，即这两个数二进制一定有一位值不同
 * 就利用这不同的一位把数组分成两个数组分别进行异或处理，这样就得到这两个数了
 */

func SingleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	xor &= -xor // 得到两个数不同的位数
	r := make([]int, 2)
	for _, num := range nums {
		if num&xor == 0 {
			r[0] ^= num
		} else {
			r[1] ^= num
		}
	}
	return r
}
