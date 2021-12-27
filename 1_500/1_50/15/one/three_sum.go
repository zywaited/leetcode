package one

import (
	"math"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var resp [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		if nums[l] < 0 && math.MinInt32-nums[l] > nums[i] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		for l < r {
			if nums[r] > -nums[l]-nums[i] {
				// 去除左边相同的数字
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				r--
			} else if nums[r] < -nums[l]-nums[i] {
				// 去除右边相同的数字
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				l++
			} else {
				resp = append(resp, []int{nums[i], nums[l], nums[r]})
				// 去除左右相同的数字
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				r--
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				l++
			}
		}
	}
	return resp
}
