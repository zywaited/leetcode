package one

func IsHappy(n int) bool {
	slow := n
	fast := next(n)
	for slow != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return slow == 1
}

func next(num int) int {
	ans := 0
	for num > 0 {
		n := num % 10
		num /= 10
		ans += n * n
	}
	return ans
}
