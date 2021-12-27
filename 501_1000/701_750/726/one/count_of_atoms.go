package one

import (
	"sort"
	"strconv"
	"strings"
)

func CountOfAtoms(formula string) string {
	type node struct {
		ele string
		num int
	}
	index := 0
	stack := []map[string]int{{}}
	fetchEle := func() string {
		currentIndex := index
		for index++; index < len(formula) && formula[index] >= 'a' && formula[index] <= 'z'; index++ {
		}
		return formula[currentIndex:index]
	}
	fetchNum := func() int {
		num := 0
		for ; index < len(formula) && formula[index] >= '0' && formula[index] <= '9'; index++ {
			num = num*10 + int(formula[index]-'0')
		}
		if num == 0 {
			num = 1
		}
		return num
	}
	for index < len(formula) {
		if formula[index] == '(' {
			stack = append(stack, map[string]int{})
			index++
			continue
		}
		if formula[index] == ')' {
			index++
			currentNum := fetchNum()
			cn := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for ele, num := range cn {
				stack[len(stack)-1][ele] += num * currentNum
			}
			continue
		}
		ele := fetchEle()
		currentNum := fetchNum()
		stack[len(stack)-1][ele] += currentNum
	}
	nums := make([]node, 0, len(stack[0]))
	for ele, num := range stack[0] {
		nums = append(nums, node{ele: ele, num: num})
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].ele < nums[j].ele
	})
	bs := strings.Builder{}
	for _, cn := range nums {
		bs.WriteString(cn.ele)
		if cn.num > 1 {
			bs.WriteString(strconv.Itoa(cn.num))
		}
	}
	return bs.String()
}
