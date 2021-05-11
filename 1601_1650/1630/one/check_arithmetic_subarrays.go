package one

func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := make([]bool, len(l))
	for i := 0; i < len(ans); i++ {
		if r[i]-l[i] == 1 {
			ans[i] = true
			continue
		}
		min := nums[l[i]]
		max := min
		ni := l[i] + 1
		for ; ni <= r[i]; ni++ {
			if nums[ni] > max {
				max = nums[ni]
				continue
			}
			if nums[ni] < min {
				min = nums[ni]
			}
		}
		if (max-min)%(r[i]-l[i]) > 0 {
			continue
		}
		diff := (max - min) / (r[i] - l[i])
		visited := make([]bool, r[i]-l[i]+1)
		for ni = l[i]; ni <= r[i]; ni++ {
			if diff == 0 {
				if nums[ni] != min {
					break
				}
				continue
			}
			if (nums[ni]-min)%diff > 0 || visited[(nums[ni]-min)/diff] {
				break
			}
			visited[(nums[ni]-min)/diff] = true
		}
		ans[i] = ni > r[i]
	}
	return ans
}
