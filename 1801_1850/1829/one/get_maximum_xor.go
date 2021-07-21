package one

func GetMaximumXor(nums []int, maximumBit int) []int {
	n := 1 << uint(maximumBit)
	ans := make([]int, 0, n)
	xor := 0
	index := 0
	for index = range nums {
		xor ^= nums[index]
	}
	mask := n - 1
	for ; index >= 0; index-- {
		ans = append(ans, (xor^mask)&mask)
		xor ^= nums[index]
	}
	return ans
}
