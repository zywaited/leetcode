package one

import "sort"

func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	v := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if v != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
