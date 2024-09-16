## [2398. 预算内的最多机器人数目](https://leetcode.cn/problems/maximum-number-of-robots-within-budget/)

### 说明
你有 n 个机器人，给你两个下标从 0 开始的整数数组 chargeTimes 和 runningCosts ，两者长度都为 n 。第 i 个机器人充电时间为 chargeTimes[i] 单位时间，花费 runningCosts[i] 单位时间运行。再给你一个整数 budget 。

运行 k 个机器人 总开销 是 max(chargeTimes) + k * sum(runningCosts) ，其中 max(chargeTimes) 是这 k 个机器人中最大充电时间，sum(runningCosts) 是这 k 个机器人的运行时间之和。

请你返回在 不超过 budget 的前提下，你 最多 可以 连续 运行的机器人数目为多少。

### 提示：
* chargeTimes.length == runningCosts.length == n
* 1 <= n <= 5 * 10^4
* 1 <= chargeTimes[i], runningCosts[i] <= 10^5
* travel.length == garbage.length - 1
* 1 <= budget <= 10^15

### 实例
#### 1
输入：chargeTimes = [3,6,1,3,4], runningCosts = [2,1,3,4,5], budget = 25
输出：3

#### 2
输入：chargeTimes = [11,12,19], runningCosts = [10,8,7], budget = 19
输出：0