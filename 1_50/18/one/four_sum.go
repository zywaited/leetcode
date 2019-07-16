package one

import "sort"

func FourSum(nums []int, target int) [][]int {
	var resp [][]int
	n := len(nums)
	if n < 4 {
		return resp
	}
	// 先排序，这样判断会简单点
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 后续不存在target，都比这个大
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 当前太小，继续下一个
		if nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// 数字相同则跳过
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1
			// 去除右边相同数字
			down := func() {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
			// 去除左边相同数字
			up := func() {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
			for l < r {
				// 太大
				if nums[i]+nums[j]+nums[l]+nums[r] > target {
					down()
					continue
				}
				// 太小
				if nums[i]+nums[j]+nums[l]+nums[r] < target {
					up()
					continue
				}
				resp = append(resp, []int{nums[i], nums[j], nums[l], nums[r]})
				up()
				down()
			}
		}
	}
	return resp
}
