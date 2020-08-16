package one

func Multiply(num1 string, num2 string) string {
	// 判断是否是0
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	nums := make([]byte, 0)
	ceil := byte(0)
	offset := 0
	for i := len(num2) - 1; i >= 0; i-- {
		to := offset
		offset++
		an := num2[i] - '0'
		for j := len(num1) - 1; j >= 0; j-- {
			if to == len(nums) {
				nums = append(nums, 0)
			}
			bn := num1[j] - '0'
			cn := an * bn
			on := nums[to] + ceil + cn%10
			nums[to] = on % 10
			ceil = cn/10 + on/10
			to++
		}
		if to == len(nums) {
			nums = append(nums, 0)
		}
		nums[to] += ceil
		ceil = 0
	}
	ans := make([]byte, 0, len(nums))
	i := len(nums) - 1
	for ; i >= 0; i-- {
		if nums[i] > 0 {
			break
		}
	}
	for ; i >= 0; i-- {
		ans = append(ans, nums[i]+'0')
	}
	return string(ans)
}
