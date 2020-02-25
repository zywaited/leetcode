package one

import "unsafe"

// 数学逻辑验证
// 1: 数字和整除3，从小到大返回所有的数字
// 2: 余数为1，则去掉最小取余为1的一个数，如果没有就去掉取余为2的两个数
// 3: 余数为2，则去掉最小取余为2的一个数，如果没有就去掉取余为1的两个数
func LargestMultipleOfThree(digits []int) string {
	nc := make([]int, 10) // 数字个数
	mc := make([]int, 3)  // 取余对应的个数
	sum := 0
	for _, num := range digits {
		nc[num]++
		mc[num%3]++
		sum += num
	}
	m := 0  // 取余
	ms := 0 // 去掉的数量
	switch sum % 3 {
	case 1:
		if mc[1] > 0 {
			m = 1
			ms = 1
		} else {
			m = 2
			ms = 2
		}
	case 2:
		if mc[2] > 0 {
			m = 2
			ms = 1
		} else {
			m = 1
			ms = 2
		}
	}
	bs := make([]byte, len(digits)-ms)
	bl := len(bs)
	for i := 0; i <= 9; i++ {
		for ; nc[i] > 0; nc[i]-- {
			if ms > 0 && i%3 == m {
				ms--
				continue
			}
			bl--
			bs[bl] = byte(i) + '0'
		}
	}
	// 无数据
	if bl == len(bs) {
		return ""
	}
	// 数字是0
	if bs[0] == '0' {
		bs = bs[:1]
	}
	return *(*string)(unsafe.Pointer(&bs))
}
