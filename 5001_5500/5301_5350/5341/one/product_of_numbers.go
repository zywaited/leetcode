package one

type ProductOfNumbers struct {
	pm []int
}

func Constructor() ProductOfNumbers {
	pn := ProductOfNumbers{pm: make([]int, 0)}
	pn.pm = append(pn.pm, 1) // 默认加一个1
	return pn
}

func (pn *ProductOfNumbers) Add(num int) {
	if num == 0 {
		// 特殊处理，还原数据
		pn.pm = pn.pm[:1]
		return
	}
	pn.pm = append(pn.pm, pn.pm[len(pn.pm)-1]*num)
}

func (pn *ProductOfNumbers) GetProduct(k int) int {
	if len(pn.pm) <= k {
		// 代表出现0值
		return 0
	}
	return pn.pm[len(pn.pm)-1] / pn.pm[len(pn.pm)-k-1]
}
