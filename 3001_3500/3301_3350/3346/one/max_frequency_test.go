package one

import (
	"fmt"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	fmt.Println(maxFrequencyV2([]int{1, 4, 5}, 1, 2))
	fmt.Println(maxFrequencyV2([]int{1}, 1, 2))
	fmt.Println(maxFrequencyV2([]int{5, 11, 20, 20}, 5, 1))
	fmt.Println(maxFrequencyV2([]int{88, 53}, 5, 1))
}
