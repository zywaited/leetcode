package one

import "sort"

func FindItinerary(tickets [][]string) []string {
	tm := make(map[string][]string, len(tickets))
	for _, ticket := range tickets {
		tm[ticket[0]] = append(tm[ticket[0]], ticket[1])
	}
	for index := range tm {
		sort.Strings(tm[index])
	}
	ans := make([]string, 0, len(tm))
	var dfs func(string)
	dfs = func(ticket string) {
		for {
			if len(tm[ticket]) == 0 {
				break
			}
			nextTicket := tm[ticket][0]
			tm[ticket] = tm[ticket][1:]
			dfs(nextTicket)
		}
		ans = append(ans, ticket)
	}
	dfs("JFK")
	// 交换位置
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
