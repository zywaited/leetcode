package one

import (
	"container/list"
	"strings"
)

func GetPermutation(n int, k int) string {
	// 每个数字开头都有(n-1)!中排列
	s := 1
	r := list.New()
	for i := 1; i < n; i++ {
		r.PushBack(i)
		s *= i
	}
	r.PushBack(n)
	rs := strings.Builder{}
	for n > 0 {
		os := k / s
		k %= s
		if k == 0 {
			os--
		}
		i := 0
		e := r.Front()
		for {
			if i != os {
				e = e.Next()
				i++
				continue
			}
			break
		}
		rs.WriteByte(byte(e.Value.(int)) + '0')
		r.Remove(e)
		if k == 0 {
			// 退出
			break
		}
		n--
		if n > 0 {
			s /= n
		}
	}
	// k == 0 时代表就是该数字开头的最后一个
	// 也就是后面数字倒序
	e := r.Back()
	for e != nil {
		rs.WriteByte(byte(e.Value.(int)) + '0')
		e = e.Prev()
	}
	return rs.String()
}
