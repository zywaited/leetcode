package one

import "sort"

func bc(candidates []int, index, target int, cr []int, resp *[][]int) {
	cl := len(candidates)
	for i := index; i < cl; i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		// 排序后发现当前比目标值还大，也就直接返回不存在
		if candidates[i] > target {
			return
		}
		if candidates[i] == target {
			*resp = append(*resp, append([]int{}, append(cr, candidates[i])...))
			continue
		}
		// 数组循环结束
		if i+1 >= cl {
			return
		}
		// 继续下一次判断
		bc(candidates, i+1, target-candidates[i], append(cr, candidates[i]), resp)
	}
}

func CombinationSum(candidates []int, target int) [][]int {
	var resp [][]int
	if len(candidates) < 1 {
		return resp
	}
	// 先排序，这样不用每次回溯到最后
	sort.Ints(candidates)
	bc(candidates, 0, target, nil, &resp)
	return resp
}
