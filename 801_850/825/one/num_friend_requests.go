package one

import "sort"

func NumFriendRequests(ages []int) int {
	ans := 0
	sort.Ints(ages)
	index101 := sort.SearchInts(ages, 101)
	for index, age := range ages {
		fi := sort.SearchInts(ages, (age>>1)+8)
		si := sort.SearchInts(ages, age+1)
		if age < 100 {
			si = min(si, index101)
		}
		if si <= fi {
			continue
		}
		num := si - fi
		if fi <= index && index < si {
			num--
		}
		ans += num
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
