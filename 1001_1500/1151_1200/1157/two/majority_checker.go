package two

// 由题目注意到threshold大于子序列长度的一半
// 这样可以采用投票法求众数
type MajorityChecker struct {
	arr []int
}

func Constructor(arr []int) MajorityChecker {
	return MajorityChecker{arr: arr}
}

func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	// 找到投票后的结果
	// 如果使用map保存当前每个数字的数量，最后就不用再循环一次，这实现理论上会快一点，但实际会出现超时
	// 可能是最后Query调用次数太多，map内存回收导致的
	rs := 0
	rn := 0
	index := left
	for ; index <= right; index++ {
		if rn == 0 {
			rs = mc.arr[index]
			rn = 1
			continue
		}
		if rs == mc.arr[index] {
			rn++
			continue
		}
		rn--
	}
	// 验证结果是否满足
	rn = 0
	for index = left; index <= right; index++ {
		if rs == mc.arr[index] {
			rn++
		}
	}
	if rn >= threshold {
		return rs
	}
	return -1
}
