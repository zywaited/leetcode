package one

func StoneGameIX(stones []int) bool {
	nums := [3]int{}
	for _, stone := range stones {
		nums[stone%3]++
	}
	check := func(i, j int) bool {
		if nums[i] == 0 {
			return false
		}
		if nums[i]-1 == nums[j] {
			return false
		}
		if nums[j] < nums[i]-2 {
			return nums[0]%2 > 0
		}
		return nums[0]%2 == 0
	}
	if nums[1] == 0 && nums[2] < 3 {
		return false
	}
	if nums[1] < 3 && nums[2] == 0 {
		return false
	}
	// 1 => 1 => 2
	// {1,2} false (1 = 2)
	// 1 => 1
	// {2} 0%2 > 0
	// 1
	// {1} 0%2 = 0
	return check(1, 2) || check(2, 1)
}
