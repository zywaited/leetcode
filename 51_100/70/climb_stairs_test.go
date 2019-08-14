package _70

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/51_100/70/one"
	"leetcode/51_100/70/three"
	"leetcode/51_100/70/two"
)

func TestClimbStairs(t *testing.T) {
	for n := 1; n <= 200; n++ {
		nums := one.ClimbStairs(n)
		require.Equal(t, nums, two.ClimbStairs(n), n)
		require.Equal(t, nums, three.ClimbStairs(n), n)
	}
}
