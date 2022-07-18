package one

func MaximumSum(nums []int) int {
	nm := [82][]int{}
	for _, num := range nums {
		sum := 0
		for tn := num; tn > 0; tn = tn / 10 {
			sum += tn % 10
		}
		if len(nm[sum]) < 2 {
			nm[sum] = append(nm[sum], num)
			continue
		}
		min := 0
		if nm[sum][1] < nm[sum][0] {
			min = 1
		}
		if nm[sum][min] < num {
			nm[sum][min] = num
		}
	}
	ans := -1
	for _, cns := range nm {
		if len(cns) < 2 {
			continue
		}
		sum := cns[0] + cns[1]
		if ans < sum {
			ans = sum
		}
	}
	return ans
}
