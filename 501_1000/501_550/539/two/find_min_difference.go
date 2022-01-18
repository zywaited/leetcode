package two

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func FindMinDifference(timePoints []string) int {
	ans := math.MaxInt32
	tr := redblacktree.NewWithIntComparator()
	for _, timePoint := range timePoints {
		curr := (int(timePoint[0]-'0')*10+int(timePoint[1]-'0'))*60 + int(timePoint[3]-'0')*10 + int(timePoint[4]-'0')
		node, found := tr.Floor(curr)
		if found {
			ans = min(ans, curr-node.Value.(int))
		}
		node, found = tr.Ceiling(curr)
		if found {
			ans = min(ans, node.Value.(int)-curr)
		}
		tr.Put(curr, curr)
	}
	ans = min(ans, 24*60-tr.Right().Value.(int)+tr.Left().Value.(int))
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
