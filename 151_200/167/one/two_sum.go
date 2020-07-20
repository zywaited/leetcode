package one

func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	ans := make([]int, 2)
	for l < r {
		if numbers[l]+numbers[r] == target {
			ans[0] = l + 1
			ans[1] = r + 1
			break
		}
		if numbers[l]+numbers[r] < target {
			l++
			continue
		}
		r--
	}
	return ans
}
