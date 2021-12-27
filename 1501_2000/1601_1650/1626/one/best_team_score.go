package one

import "sort"

func BestTeamScore(scores []int, ages []int) int {
	type node struct {
		score int
		age   int
	}
	var nodes []node
	for index := range scores {
		nodes = append(nodes, node{score: scores[index], age: ages[index]})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].age == nodes[j].age {
			return nodes[i].score < nodes[j].score
		}
		return nodes[i].age < nodes[j].age
	})
	dp := make([]int, len(nodes))
	ans := 0
	for i := 0; i < len(nodes); i++ {
		dp[i] = nodes[i].score
		for j := i - 1; j >= 0; j-- {
			if nodes[j].score <= nodes[i].score {
				dp[i] = max(dp[i], dp[j]+nodes[i].score)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
