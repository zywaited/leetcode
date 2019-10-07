package one

import (
	"strconv"
	"strings"
)

type Solution struct {
	target        int
	nums          string
	buff          strings.Builder
	currOperators []string
	rs            []string
}

func AddOperators(num string, target int) []string {
	s := &Solution{
		target: target,
		nums:   num,
	}
	s.bc(0, 0, 0, 0)
	return s.rs
}

// 回溯
// prev 前一个操作的数字
// curr 当前正在操作的数字
// value 已经计算的值
func (s *Solution) bc(index, prev, curr, value int) {
	if index == len(s.nums) {
		// 已经结束了
		if value == s.target && curr == 0 {
			// 等于目标值，并且没有下一个值
			s.buff.Reset()
			for _, operator := range s.currOperators {
				s.buff.WriteString(operator)
			}
			s.rs = append(s.rs, s.buff.String())
		}
		return
	}
	// 可以耦合数字
	curr = curr*10 + int(s.nums[index]-'0')
	currRep := strconv.Itoa(curr)
	if curr > 0 {
		// 如果curr为0的话后面就不能耦合数字，因为0(\d)不合法
		s.bc(index+1, prev, curr, value)
	}
	if len(s.currOperators) == 0 {
		// 最开始没有符号
		s.currOperators = append(s.currOperators, currRep)
		s.bc(index+1, curr, 0, value+curr)
		s.currOperators = s.currOperators[:len(s.currOperators)-1]
		return
	}
	// +
	s.currOperators = append(s.currOperators, "+")
	s.currOperators = append(s.currOperators, currRep)
	s.bc(index+1, curr, 0, value+curr)
	s.currOperators = s.currOperators[:len(s.currOperators)-2]
	// -
	s.currOperators = append(s.currOperators, "-")
	s.currOperators = append(s.currOperators, currRep)
	s.bc(index+1, -curr, 0, value-curr) // 把curr置为-curr，是为了后面都以加法操作
	s.currOperators = s.currOperators[:len(s.currOperators)-2]
	// *
	s.currOperators = append(s.currOperators, "*")
	s.currOperators = append(s.currOperators, currRep)
	s.bc(index+1, prev*curr, 0, value-prev+(prev*curr))
	s.currOperators = s.currOperators[:len(s.currOperators)-2]
}
