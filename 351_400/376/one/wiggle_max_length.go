package one

func WiggleMaxLength(nums []int) int {
	n := 0
	prev := 0
	f := false
	for _, num := range nums {
		if n == 0 {
			n++
			prev = num
			continue
		}
		if prev == num {
			continue
		}
		if n == 1 {
			f = num > prev
			n++
			prev = num
			continue
		}
		if !f {
			if num < prev {
				prev = num
				continue
			}
			n++
			prev = num
			f = true
			continue
		}
		if num > prev {
			prev = num
			continue
		}
		f = false
		n++
		prev = num
	}
	return n
}
