package four

func SubarrayBitwiseORs(arr []int) int {
	maxBit := 0
	for _, num := range arr {
		maxBit |= num
	}
	nums := make(map[int]byte)
	for i := 0; i < len(arr); i++ {
		mask := arr[i]
		nums[mask] = 1
		for j := i - 1; j >= 0 && mask < maxBit; j-- {
			mask |= arr[j]
			nums[mask] = 1
		}
	}
	return len(nums)
}
