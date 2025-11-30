package one

import (
	"fmt"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	fmt.Println(maxFrequency([]int{5, 11, 20, 20}, 5, 1))
	fmt.Println(maxFrequency([]int{88, 53}, 5, 1))
}
