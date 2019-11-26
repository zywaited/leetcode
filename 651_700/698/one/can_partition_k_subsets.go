package one

import "sort"

func CanPartitionKSubsets(nums []int, k int) bool {
	// 排序后不用重复计算
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 每个子集的总和
	tr, ex := sum/k, sum%k
	if ex != 0 {
		return false
	}
	end := len(nums) - 1
	for end >= 0 && nums[end] >= tr {
		// 比目标值大，不可能有子集
		if nums[end] > tr {
			return false
		}
		// 单独成一组
		k--
		end--
	}
	gs := make([]int, k)
	var sr func(int) bool
	sr = func(end int) bool {
		// 如果结束
		if end < 0 {
			return true
		}
		// 从nums中找到能够满足target的集合(递归)
		// 前面处理完后，nums最大值比tr小
		for j := 0; j < k; j++ {
			if gs[j]+nums[end] <= tr {
				gs[j] += nums[end]
				if sr(end - 1) {
					// 分组完成
					return true
				}
				// 该数字不能使用在当前
				gs[j] -= nums[end]
			}
			// 如果最后该组没有数据
			if gs[j] == 0 {
				return false
			}
		}
		return false
	}
	return sr(end)
}
