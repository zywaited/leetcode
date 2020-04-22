package one

import "math"

func IncreasingTriplet(nums []int) bool {
	f, s := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= f {
			f = num
			continue
		}
		if num <= s {
			s = num
			continue
		}
		return true
	}
	return false
}
