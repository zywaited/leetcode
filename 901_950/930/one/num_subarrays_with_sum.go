package one

func NumSubarraysWithSum(nums []int, goal int) int {
	ans := 0
	prev := 0
	last := 0
	sum := 0
	zeroNum := 1
	for ; last < len(nums); last++ {
		sum += nums[last]
		if sum < goal {
			continue
		}
		if sum > goal {
			for ; prev <= last && sum > goal; prev++ {
				sum -= nums[prev]
			}
			zeroNum = 1
		}
		if prev > last {
			continue
		}
		for ; prev < last && nums[prev] == 0; prev++ {
			zeroNum++
		}
		ans += zeroNum
	}
	return ans
}
