package one

func PeopleAwareOfSecret(n int, delay int, forget int) int {
	mod := int(1e9 + 7)
	type secret struct {
		day int
		num int
	}
	canTalk := make([]secret, 0, n)
	canTalkSum := 0
	onlyKnow := make([]secret, 0, n)
	onlyKnowSum := 1
	onlyKnow = append(onlyKnow, secret{day: 1, num: 1})
	for i := 2; i <= n; i++ {
		// 去除已经忘记秘密的人
		for canTalkSum > 0 && canTalk[0].day+forget <= i {
			canTalkSum = (canTalkSum - canTalk[0].num + mod) % mod
			canTalk = canTalk[1:]
		}
		// 可以说秘密的人
		for onlyKnowSum > 0 && onlyKnow[0].day+delay <= i {
			canTalk = append(canTalk, onlyKnow[0])
			canTalkSum = (canTalkSum + onlyKnow[0].num) % mod
			onlyKnowSum = (onlyKnowSum - onlyKnow[0].num + mod) % mod
			onlyKnow = onlyKnow[1:]
		}
		// 新知道秘密的人
		onlyKnow = append(onlyKnow, secret{day: i, num: canTalkSum})
		onlyKnowSum = (onlyKnowSum + canTalkSum) % mod
	}
	return (canTalkSum + onlyKnowSum) % mod
}
