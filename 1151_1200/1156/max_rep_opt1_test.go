package _1156

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1151_1200/1156/one"
)

func TestMaxRepOpt1(t *testing.T) {
	require.Equal(t, 3, one.MaxRepOpt1("ababa"))
	require.Equal(t, 8, one.MaxRepOpt1("aaaaaaaa"))
	require.Equal(t, 6, one.MaxRepOpt1("aaabaaa"))
	require.Equal(t, 4, one.MaxRepOpt1("aaabbaaa"))
	require.Equal(t, 1, one.MaxRepOpt1("abcdef"))
	require.Equal(t, 4, one.MaxRepOpt1("abaaa"))
}
