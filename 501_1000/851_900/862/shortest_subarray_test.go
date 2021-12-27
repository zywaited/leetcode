package _862

import (
	"leetcode/501_1000/851_900/862/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortestSubArray(t *testing.T) {
	require.EqualValues(t, 24, one.ShortestSubarray([]int{18393, 13015, 39877, -46990, 84110, -16361, -42889, -10085, 46644, 91545, 71230, 48090, 34489, 2788, 56496, -32528, 77455, -44282, 80504, 77949, 70, 74450, -4005, 82184, -19401, 49038, -10000, 31028, 26603, 62302, 76071, 73298, -20008, -12421, -11904, -8055, 50871, 20857, 56174, 88271, 37380, 56974, 85085, -29377, -39430, 86935, -42349, -12415, -21752, 95087}, 917790))
	require.EqualValues(t, 9, one.ShortestSubarray([]int{-36, 10, -28, -42, 17, 83, 87, 79, 51, -26, 33, 53, 63, 61, 76, 34, 57, 68, 1, -30}, 484))
	require.EqualValues(t, 2, one.ShortestSubarray([]int{-34, 37, 51, 3, -12, -50, 51, 100}, 151))
}
