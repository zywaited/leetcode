package one

import "sort"

func BraceExpansionII(expression string) []string {
	cm, _, _ := braceExpansion(expression, 0)
	ans := make([]string, 0, len(cm))
	for s := range cm {
		ans = append(ans, s)
	}
	sort.Strings(ans)
	return ans
}

func braceExpansion(expression string, index int) (cm map[string]bool, pm map[string]bool, nIndex int) {
	cm = make(map[string]bool) // 最终数据
	pm = make(map[string]bool) // 未合并数据
	for index < len(expression) {
		switch expression[index] {
		case '{':
			ncm, npm, nextIndex := braceExpansion(expression, index+1)
			index = nextIndex
			if len(pm) == 0 {
				pm = ncm
				break
			}
			for s := range pm {
				for ns := range ncm {
					npm[s+ns] = true
				}
			}
			pm = npm
		case '}':
			merge(cm, pm)
			nIndex = index + 1
			return
		case ',':
			// 合并
			merge(cm, pm)
			index++
		default:
			pIndex := index
			for ; index < len(expression) && expression[index] >= 'a' && expression[index] <= 'z'; index++ {
			}
			if len(pm) == 0 {
				pm[expression[pIndex:index]] = true
				break
			}
			npm := make(map[string]bool, len(pm))
			for s := range pm {
				npm[s+expression[pIndex:index]] = true
			}
			pm = npm
		}
	}
	merge(cm, pm)
	nIndex = len(expression)
	return
}

func merge(cm, pm map[string]bool) {
	for s := range pm {
		cm[s] = true
		delete(pm, s)
	}
}
