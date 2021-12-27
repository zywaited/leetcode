package one

func Rotate(nums []int, k int) {
	num := 0
	for i := 0; i < k && num < len(nums); i++ {
		j := i
		tmp := nums[j]
		for (j+k)%len(nums) != i {
			j = (j + k) % len(nums)
			nums[j], tmp = tmp, nums[j]
			num++
		}
		nums[i] = tmp
		num++
	}
}
