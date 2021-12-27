package one

import "sort"

var (
	// 返回结果
	resp [][]int
	// 数组长度
	cl = 0
	// 原始数据
	cds []int
)

func bc(index int, target int, cr []int) {
	if target == 0 {
		resp = append(resp, append([]int{}, cr...))
		return
	}

	for i := index; i < cl; i++ {
		if target < cds[i] {
			return
		}
		// 无限次使用当前数字
		cr = append(cr, cds[i])
		bc(i, target-cds[i], cr)
		cr = cr[:len(cr)-1]
	}
}

func CombinationSum(candidates []int, target int) [][]int {
	// 先排序， 不然每次都要遍历所有
	sort.Ints(candidates)
	cds = candidates
	cl = len(cds)
	resp = nil
	var cr []int
	bc(0, target, cr)
	return resp
}
