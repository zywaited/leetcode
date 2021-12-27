package one

import "sort"

func FindBestValue(arr []int, target int) int {
	sort.Ints(arr)
	// 前缀和
	ps := make([]int, len(arr)+1)
	for index, num := range arr {
		ps[index+1] = ps[index] + num
	}
	// 二分搜索比当前索引大的第一个索引
	bs := func(s, num int) int {
		if arr[s] > num {
			return s
		}
		e := len(arr) - 1
		i := s + 1
		for s <= e {
			m := s + (e-s)>>1
			if arr[m] <= num {
				s = m + 1
				continue
			}
			i = m
			e = m - 1
		}
		return i
	}
	// 起始值
	num := target / len(arr)
	ans := -1
	fs := 0
	fi := 0
	for {
		index := bs(fi, num)
		sum := ps[index] + num*(len(arr)-index)
		if sum == target {
			return num
		}
		if sum < target {
			ans = num
			fs = sum
			fi = index
			if index == len(arr) {
				// 已经是最大值
				break
			}
			num = arr[index]
			continue
		}
		if ans == -1 || ans == num {
			ans = num
			break
		}
		// 这里代表target在中间
		for nn := ans + 1; nn <= num; nn++ {
			sum = ps[fi] + nn*(len(arr)-fi)
			if abs(sum-target) < abs(fs-target) {
				ans = nn
				fs = sum
				continue
			}
			break
		}
		break
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
