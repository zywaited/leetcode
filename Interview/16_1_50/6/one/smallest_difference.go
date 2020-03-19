package one

import (
	"math"
	"sort"
)

// NlogN
func SmallestDifference(a []int, b []int) int {
	// 先对a进行排序，然后对b在a中二分查找最接近的数字
	sort.Ints(a)
	ans := math.MaxInt32
	for _, num := range b {
		num = find(a, num)
		if num < ans {
			ans = num
		}
	}
	return ans
}

func find(nums []int, target int) int {
	s, e := 0, len(nums)-1
	rs := math.MaxInt32
	for s <= e {
		m := s + (e-s)>>1
		if nums[m] == target {
			rs = 0
			break
		}
		if valid(nums[m], target) {
			rs = min(rs, abs(nums[m]-target))
		}
		if nums[m] < target {
			s = m + 1
			continue
		}
		e = m - 1
	}
	return rs
}

// 判断是否过界
func valid(f, s int) bool {
	if f >= 0 && s >= 0 || f <= 0 && s <= 0 {
		return true
	}
	if f >= 0 && math.MinInt32+f <= s {
		return true
	}
	return math.MinInt32+s <= f
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
