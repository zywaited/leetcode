package one

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
			continue
		}
		r = m - 1
	}
	return l
}
