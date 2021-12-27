package one

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	if len(nums) > 0 {
		// 先转成字符串
		c := make(customer, len(nums))
		for i, num := range nums {
			c[i] = strconv.Itoa(num)
		}
		sort.Sort(c)
		// 首位不能为0
		if c[0] == "0" {
			return "0"
		}
		s := strings.Builder{}
		for _, m := range c {
			s.WriteString(m)
		}
		return s.String()
	}
	return "0"
}

type customer []string

func (c customer) Len() int {
	return len(c)
}

func (c customer) Less(i, j int) bool {
	// 比较两个连接那个大
	f, s := c[i]+c[j], c[j]+c[i]
	i = 0
	for {
		if i == len(f) && i == len(s) || f[i] > s[i] {
			return true
		}
		if f[i] < s[i] {
			return false
		}
		i++
	}
}

func (c customer) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
