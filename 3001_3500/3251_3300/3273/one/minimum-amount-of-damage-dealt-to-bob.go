package one

func minDamage(power int, damage []int, health []int) int64 {
	type enemy struct {
		damage   int
		survival int
	}
	enemies := make([]enemy, len(damage))
	damages := int64(0)
	for i := 0; i < len(damage); i++ {
		enemies[i] = enemy{damage: damage[i], survival: (health[i] + power - 1) / power}
		damages += int64(damage[i])
	}
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].damage*enemies[j].survival > enemies[j].damage*enemies[i].survival
	})
	total := int64(0)
	for i := 0; i < len(enemies); i++ {
		total += damages * int64(enemies[i].survival)
		damages -= int64(enemies[i].damage)
	}
	return total
}
