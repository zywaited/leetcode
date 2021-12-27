package _238

import (
	"leetcode/1_500/201_250/238/one"
	"leetcode/1_500/201_250/238/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductExceptSelf(t *testing.T) {
	require.Equal(
		t,
		one.ProductExceptSelf([]int{1, 2, 3, 4}),
		two.ProductExceptSelf([]int{1, 2, 3, 4}),
	)

	require.Equal(
		t,
		one.ProductExceptSelf([]int{1, 2, 3, 4, 5, 6, 7, 8}),
		two.ProductExceptSelf([]int{1, 2, 3, 4, 5, 6, 7, 8}),
	)
}
