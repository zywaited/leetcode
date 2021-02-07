## [1423. 可获得的最大点数](https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/)

### 说明
几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。

每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。

你的点数就是你拿到手中的所有卡牌的点数之和。

给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。

* 1 <= cardPoints.length <= 10^5
* 1 <= cardPoints[i] <= 10^4
* 1 <= k <= cardPoints.length

### 实例
#### 1
输入：cardPoints = [1,2,3,4,5,6,1], k = 3
输出：12

#### 2
输入：cardPoints = [2,2,2], k = 2
输出：4