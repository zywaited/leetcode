## [3273. 对 Bob 造成的最少伤害](https://leetcode.cn/problems/minimum-amount-of-damage-dealt-to-bob/)

### 说明
给你一个整数 power 和两个整数数组 damage 和 health ，两个数组的长度都为 n 。

Bob 有 n 个敌人，如果第 i 个敌人还活着（也就是健康值 health[i] > 0 的时候），每秒钟会对 Bob 造成 damage[i] 点 伤害。

每一秒中，在敌人对 Bob 造成伤害 之后 ，Bob 会选择 一个 还活着的敌人进行攻击，该敌人的健康值减少 power 。

请你返回 Bob 将 所有 n 个敌人都消灭之前，最少 会受到多少伤害。

### 提示：
* 1 <= power <= 10^4
* 1 <= n == damage.length == health.length <= 10^5
* 1 <= damage[i], health[i] <= 10^4

### 实例
#### 1
输入：power = 4, damage = [1,2,3,4], health = [4,5,6,8]
输出：39

#### 2
输入：power = 1, damage = [1,1,1,1], health = [1,2,3,4]
输出：20