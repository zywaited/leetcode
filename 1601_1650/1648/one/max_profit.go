package one

import "sort"

func MaxProfit(inventory []int, orders int) int {
	mod := int(1e9 + 7)
	ans := 0
	sort.Ints(inventory)
	times := 0
	num := 0
	buy := orders
	remainder := 0
	multi := 0
	for index := len(inventory) - 1; index >= 0 && buy > 0; index-- {
		multi++
		num = inventory[index]
		times = inventory[index]
		if index > 0 {
			times = inventory[index] - inventory[index-1]
		}
		remainder = 0
		if buy < times*multi {
			times = buy / multi
			remainder = buy % multi
		}
		buy -= times * multi
		num -= times - 1
		ans = (ans + num*times*multi) % mod
		ans = (ans + (times-1)*times/2*multi) % mod
		if remainder > 0 {
			ans = (ans + (num-1)*remainder) % mod
			break
		}
	}
	return ans
}
