package one

import "strconv"

func CountAndSay(n int) string {
	prev := "1"
	bs := []byte{}
	for i := 1; i < n; i++ {
		bs = bs[:0]
		num := 0
		j := 0
		for j < len(prev) {
			num = 1
			for j++; j < len(prev) && prev[j] == prev[j-1]; j++ {
				num++
			}
			bs = append(bs, strconv.Itoa(num)...)
			bs = append(bs, prev[j-1])
		}
		prev = string(bs)
	}
	return prev
}
