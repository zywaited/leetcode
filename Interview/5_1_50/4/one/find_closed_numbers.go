package one

func FindClosedNumbers(num int) []int {
	// 比当前小的最接近的数就是从后往前的1后面一位是0的右移一位(后面的1全部靠左)
	// 比当前大的最接近的数就是从后往前的1前面一位是0的左移一位(后面的1全部靠右)
	find := 0
	ans := []int{-1, -1}
	on := uint(0) // 去除数中1的个数
	for num > 0 {
		pn := num & (-num)
		num &= num - 1
		if pn&(1<<30) == 0 && num&(pn<<1) == 0 && ans[0] == -1 {
			ans[0] = num | (pn << 1) | (1<<on - 1)
			find++
		}
		if pn > 1 && num&(pn>>1) == 0 && ans[1] == -1 {
			ans[1] = num | (pn >> 1) | ((pn>>1 - 1) ^ (pn>>1-1)>>on)
			find++
		}
		if find == 2 {
			break
		}
		on++
	}
	return ans
}
