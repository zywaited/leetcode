package _212

import (
	"testing"

	"leetcode/201_250/212/one"

	"github.com/stretchr/testify/require"
)

func TestFindWords(t *testing.T) {
	words := []string{"oath", "pea", "eat", "rain"}
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	outs := []string{"eat", "oath"}
	finds := one.FindWords(board, words)
	require.Len(t, finds, len(outs))
	for _, word := range outs {
		require.Contains(t, finds, word)
	}
}
