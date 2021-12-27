package _1156

import (
	"leetcode/1001_1500/1151_1200/1156/one"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxRepOpt1(t *testing.T) {
	require.Equal(t, 3, one.MaxRepOpt1("ababa"))
	require.Equal(t, 8, one.MaxRepOpt1("aaaaaaaa"))
	require.Equal(t, 6, one.MaxRepOpt1("aaabaaa"))
	require.Equal(t, 4, one.MaxRepOpt1("aaabbaaa"))
	require.Equal(t, 1, one.MaxRepOpt1("abcdef"))
	require.Equal(t, 4, one.MaxRepOpt1("abaaa"))
}
