package two

import (
	"math/bits"
	"strconv"
)

func ReadBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if bits.OnesCount(uint(i))+bits.OnesCount(uint(j)) == turnedOn {
				ans = append(ans, strconv.Itoa(i)+":"+strconv.Itoa(j/10)+strconv.Itoa(j%10))
			}
		}
	}
	return ans
}
