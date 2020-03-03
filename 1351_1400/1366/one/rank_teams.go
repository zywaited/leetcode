package one

import (
	"sort"
	"strings"
)

func RankTeams(votes []string) string {
	// 统计所有排名数量，然后依次对比
	// 本来可以直接定义数组26，不过不想进行过多无意义排序(当长度很少的时候, 所以用个map记录索引)
	//scores := [26][26]int{}
	scores := make([]struct {
		team  byte
		score []int
	}, len(votes[0]))
	// byte index
	indexes := make(map[byte]int, len(votes[0]))
	// 初始化
	for i := range votes[0] {
		indexes[votes[0][i]] = i
		scores[i].team = votes[0][i]
		scores[i].score = make([]int, len(votes[0]))
	}
	// 统计排名
	for _, vote := range votes {
		for i := range vote {
			scores[indexes[vote[i]]].score[i]++
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		// 是否i的排名比j高
		for k, num := range scores[i].score {
			if num == scores[j].score[k] {
				continue
			}
			return num > scores[j].score[k]
		}
		// 是否i的字母小
		return scores[i].team < scores[j].team
	})
	ans := strings.Builder{}
	for _, score := range scores {
		ans.WriteByte(score.team)
	}
	return ans.String()
}
